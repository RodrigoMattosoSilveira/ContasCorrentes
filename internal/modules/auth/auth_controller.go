package auth

import (
	"log"
	"os"
	"time"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/modules/people"
	constants "github.com/RodrigoMattosoSilveira/ContasCorrentes/constants"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	db    *gorm.DB
	store *session.Store
}

type LoginMessage struct {
	Message string
}
		
func NewAuthController(db *gorm.DB, store *session.Store) *AuthController {
	return &AuthController{db: db, store: store}
}

func (ac *AuthController) ShowLoginForm(c *fiber.Ctx) error {
	return c.Render("pages/auth/login", fiber.Map{
		"Title":      "Login",
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"PersonFirst":   c.Locals("PersonFirst"),
	}, "layouts/base")
}

func (ac *AuthController) HandleLogin(c *fiber.Ctx) error {
	messageLogin := LoginMessage{}
	type LoginRequest struct {
		Email string `form:"email"`
		Password string `form:"password"`
	}
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		// return c.Status(400).SendString("Invalid request")
		messageLogin.Message = "Invalid Request"
		return c.Render("partials/auth/authMessage", messageLogin)
	}

	var person people.Person
	if err := ac.db.Where("email = ?", req.Email).First(&person).Error; err != nil {
		// return c.Status(401).SendString("Invalid email")
		messageLogin.Message = "Invalid user name"
		return c.Render("partials/auth/authMessage", messageLogin)
	}

	if !CheckPasswordHash(person.Password, req.Password) {
		// return c.Status(401).SendString("Invalid password")
		messageLogin.Message = "Invalid user password"
		return c.Render("partials/auth/authMessage", messageLogin)
	}	

	sess, err := ac.store.Get(c)
	if err != nil {
		// return c.Status(500).SendString("Session error")
		messageLogin.Message = "Session error"
		return c.Render("partials/auth/authMessage", messageLogin)
	}

	sess.Set("PersonId", person.ID)
	sess.Set("PersonName", person.Name)

	if err := sess.Save(); err != nil {
		log.Printf("ERROR: Failed to save session: %v", err)
		// return c.Status(500).SendString("Failed to save session")
		messageLogin.Message = "Failed to save session"
		return c.Render("partials/auth/authMessage", messageLogin)
	}
	c.Set("HX-Redirect", "/profile")

    // Generate JWT token
    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": person.Name,
		"user_id": person.Email,
		"role": "Associate",
   		"iss": "ContasCorrentes",     
        "exp": 60 * 60 * 6 + time.Now().Unix(), // 6 hours from now
    	"iat": time.Now().Unix(),
    })

    secretKey := os.Getenv("JWT_KEY")
    token, err := claims.SignedString([]byte(secretKey))
    if err != nil {
		// return c.SendStatus(fiber.StatusInternalServerError)
		messageLogin.Message = "Fiber Internal Server Error"
		return c.Render("partials/auth/authMessage", messageLogin)
    }

	// Create jwt cookie
	cookie := new(fiber.Cookie)
	cookie.Name = constants.COOKIE_NAME
	cookie.Value =  token
	cookie.MaxAge = 1000*60*60*6 // 6 hours
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.SameSite = "Secure"
	c.Cookie(cookie)

	return c.SendStatus(200)
}

func (ac *AuthController) ShowProfile(c *fiber.Ctx) error {
	sess, err := ac.store.Get(c)
	if err != nil {
		return c.Status(500).SendString("Session error")
	}

	personID := sess.Get("PersonId")
	if personID == nil {
		return c.Redirect("/login")
	}

	return c.Render("pages/auth/profile", fiber.Map{
		"Title":      "Your Profile",
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"PersonFirst":   sess.Get("PersonFirst"),
		"PersonLast":    sess.Get("PersonLast"),
	}, "layouts/base")
}

func (ac *AuthController) AuthorizationError(c *fiber.Ctx) error {
	// sess, err := ac.store.Get(c)
	// if err != nil {
	// 	return c.Status(500).SendString("Session error")
	// }
	log.Println("AuthorizationError, Query Params:", c.OriginalURL()) // Log the full URL with query params

	return c.Render("partials/auth/authorizationError", nil)
}

func (ac *AuthController) HandleLogout(c *fiber.Ctx) error {
	sess, err := ac.store.Get(c)
	if err != nil {
		return c.Status(500).SendString("Session error")
	}

	if err := sess.Destroy(); err != nil {
		log.Printf("ERROR: Failed to destroy session: %v", err)
		return c.Status(500).SendString("Failed to destroy session")
	}

	return c.Redirect("/")
}

func CheckPasswordHash(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
