package middleware

import (
	"errors"
	"log"
	"os"
	"time"

	constants "github.com/RodrigoMattosoSilveira/ContasCorrentes/constants"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

/*
	REQUIRED(Any middleware must have this)

	For every middleware we need a config.
	In config we also need to define a function which allows us to skip the middleware if return true.
	By convention it should be named as "Filter" but any other name will work too.
*/
type Config struct {
	// when returned true, our middleware is skipped
	Filter         func(c *fiber.Ctx) bool

	// function to run when there is error decoding jwt
	Unauthorized  func(c *fiber.Ctx) error

	// function to decode our jwt token
	Decode       func(c *fiber.Ctx) (*jwt.MapClaims, error)

	// set jwt secret
	Secret       string

	// set jwt expiry in seconds
	Expiry       int64
}

/*
	Middleware specific

	Our middleware's config default values if not passed
*/
var ConfigDefault = Config{
	Filter:       nil,
	Decode:       nil,
	Unauthorized: nil,
	Secret:       "a_very_weak_secret",
	Expiry:       60 * 60 * 6, // 6 hours
}

/*
	Middleware specific

	Function for generating default config
*/
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values if not passed
	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	// Set default secret if not passed
	if cfg.Secret == "" {
		cfg.Secret = ConfigDefault.Secret
	}

	// Set default expiry if not passed
	if cfg.Expiry == 0 {
		cfg.Expiry = ConfigDefault.Expiry
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = Unauthorized
	}

    // this is the main jwt decode function of our middleware
	if cfg.Decode == nil {
		// Set default Decode function if not passed
		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {

			// Retrieve JWT token from cookie
			cookie := c.Cookies(constants.COOKIE_NAME) // Replace "jwt" with your actual cookie name if different

			// Parse JWT token with claims
			secretKey := os.Getenv("JWT_KEY")
			// token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 	return []byte(secretKey), nil
			// })
			token, err := jwt.Parse(cookie, func(token *jwt.Token) (any, error) {
				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(secretKey), nil
			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))			// Handle token parsing errors
			if err != nil {
				return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("invalid token")
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				return nil, errors.New("jwt is expired")
			}

			return &claims, nil
		}
	}

	// Set default Unauthorized if not passed
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}
/*
	Middleware specific
	
	Function to generate a jwt token
*/
func Encode(claims *jwt.MapClaims, expiryAfter int64) (string, error) {

	// setting default expiryAfter
	if expiryAfter == 0 {
		expiryAfter = ConfigDefault.Expiry
	}
        
        // or you can use time.Now().Add(time.Second * time.Duration(expiryAfter)).UTC().Unix()
	(*claims)["exp"] = time.Now().UTC().Unix() + expiryAfter

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// our signed jwt token string
	secretKey := os.Getenv("JWT_KEY")
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", errors.New("error creating a token")
	}

	return signedToken, nil
}

/*
    REQUIRED(Any middleware must have this)

	Our main middleware function used to initialize our middleware.
	By convention we name it "New" but any other name will work too.
*/
func New(config Config) fiber.Handler {

	// For setting default config
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Filter returns true
		if cfg.Filter != nil && cfg.Filter(c) {
			log.Println("Midddle was skipped")
			return c.Next()
		}
		log.Println("Midddle was run")

		claims, err := cfg.Decode(c)

		if err == nil {
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}

func Unauthorized(ctx *fiber.Ctx) error {
	log.Println("Unauthorized, Query Params:", ctx.OriginalURL()) // Log the full URL with query params
	ctx.Set("HX-Trigger", "unauthorized")
	return ctx.SendString("Triggering with HX-Trigger")
}	