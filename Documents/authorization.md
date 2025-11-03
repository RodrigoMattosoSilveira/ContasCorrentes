# Introduction
I implemented the `Authentication` and `Authorizarion` do be as simple as possible, using the best practices I discerned from the vast material I read. The goal of the `Authentication` flow is to ensure that a guest requesting permision to use the system is who they claim to be. The goal of the `Authorization` flow is to ensure that a `Person` using the system has been `Authenticated` and is authorized to access and operate on resources prior to doing it. 

# Authentication
The `Authentication` flow has two paths, i. `signup`, when a guest requests an account, ii. `signin`, when a guest with an account requests permision to use the system.

I achieve it by requesting a user identification (their email, which has to be unique) and a password (which has to be secure) when they sign up to use the system. When persisting a new Person Entity, I encript the password in order to ensure that no one will ever see it again, therefore preserving the person's identity. 

When a guest attempts to login to gain access to the system, I prompt them for their email and password. I use the email to ensure that a person with that email exists in the People database. I encrypt the password provided in the login form and compare it with the encrypted password in retrieved person's record. If they match I assume that the guest is the person who they claim to be, and allow them to use the system.

## Signup
The system home page includes links for a guest to `signup` or `signin`. When a guest requests an account, the system displays a form request their email and password, among other data, then routes the request to a function that extracts them into a `Person` record and uses it to complete adding the new person's account.

``` go
/* people-controller*/
func (uc *PeopleController) AddPerson(c *fiber.Ctx) error {
	/*
	...
	 */
	hashedPassword, err := HashPassword(person.Password)
	if err != nil {
		return c.Status(400).SendString("Failed to hash password:")
	}
	person.Password = hashedPassword

	/*
	...
	 */

}
```
**Notes**:
Of interest on the account adding flow is the password encryption; I;
- use `Golang Crypto` to encript the password provided by the guest; 
- persist their `Person` entity using the encrypted password;
- no one will be able to see the actual password provided by the guest, not even the system!
  
## Password Verification
**TBD**

## Signin
 When a guest requests an authentication, the system displays a form with their email and password, then routes the request to a function that extracts them into a `Login` structure and uses it to complete the authentication request;

``` go 
/* auth_controller.go */
func (ac *AuthController) HandleLogin(c *fiber.Ctx) error {
	/*
	...
	 */
	// Retrieve the data in the log in form
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

	// Use the login email to rerive the guest's Person record
	var person people.Person
	if err := ac.db.Where("email = ?", req.Email).First(&person).Error; err != nil {
		// return c.Status(401).SendString("Invalid email")
		messageLogin.Message = "Invalid user name"
		return c.Render("partials/auth/authMessage", messageLogin)
	}

	// If the person record exists, compare the passwords
	if !CheckPasswordHash(person.Password, req.Password) {
		// return c.Status(401).SendString("Invalid password")
		messageLogin.Message = "Invalid user password"
		return c.Render("partials/auth/authMessage", messageLogin)
	}	
	/*
	 ...
	*/
}
```
**Notes**:
  When a guest requests authentication, I;
  - collect their form data (email and password);
  - use the email to retrieve their `Person` record;
  - if there is no record with the email, I return a message with a `401 Unauthorized status code`  HTTP Status
  - encrypt the provided password, and compare with the encrypted `Person` password;
  - if they match, I return a `200 ok` HTTP Status, otherwise a `401 Unauthorized status code`  HTTP Status

## Forgot Password
**TBD**

# Authorization
The goal of the `Authorization` flow is to ensure that a `Person` using the system is `Authenticated` and can has the permission to access and operate on resources at any point in time. I used a `Role`, part of a hierarqui, persisted in the `Person` record, to ensure a person is authorized to perform an operation. For now we will have the following roles:
- Person - Can see and update their own data;
- Operator - Can perform  basic personnel, planning, mercantile, and well operations
- Administrator -  Can perform  all personnel, planning, mercantile, and well operations
- System - Can do anything, including adding a new system administrator

When a `guest` signs in I define a `token` including claims, such as expiration date and personId, about the authenticated `person`; I encrypt the token and include it as a `cookie` in the `Authentication Response`; from then on, until the `person` logs out, the Client Agent will include the cookie with the encripted token the all the person's requests. 

``` go
    /*
	... auth_controller.go / func (ac *AuthController) HandleLogin(c *fiber.Ctx) error 
	 */
    // Generate JWT token
	// TODO add logic to only pass the Person.ID here
    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": person.ID,
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
   /*
	...
	 */
```

I defined `middleware` logic that intercepts all`HTTP Requests`, decrypts and validates a `cookie` therein, into a `JWT Token`; if it is decrypted correctly, it extracts the `expiraton date` from the JWT Token and ensures it has not expired; next, I extract the personId, read the Person record, the Person Role, save the personId and personRole in the Session, and call the Next middleware function.

``` go
/*
 ... authorization.go / func configDefault(config ...Config) Config
 */
			// Retrieve JWT token from cookie
			cookie := c.Cookies(constants.COOKIE_NAME) // Replace "jwt" with your actual cookie name if different
			if cookie == "" {
				c.Locals("AuthenticationErrorText", "Authentication cookie not found. Please log in again.")
				c.Locals("AuthenticationError", "Authentication cookie not found. Please log in again.")
				return nil, errors.New("cookie not found")
			}

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
				c.Locals("AuthenticationErrorText", "Invalid authentication token. Please log in again.")
				return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}

			// TODO add logic to retrieve the Person.Role here
			claims, ok := token.Claims.(jwt.MapClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("invalid token")
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				c.Locals("AuthenticationErrorText", "Authentication cookieexpired. Please log in again.")
				return nil, errors.New("jwt is expired")
			}

			return &claims, nil

```
Next, I configure the routes that require authorization; the `home` and `/login` routes are examples of routes that do not require authorizarion.
``` go
/*
 ... person_routes.go / func RegisterRoutes(router fiber.Router, db *gorm.DB)
 */
 func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	service := NewPeopleService(db)
	controller := NewPeopleController(service)

	group := router.Group("/people", middleware.New(middleware.Config{}))

	group.Get("/", controller.ListPeople)
	group.Post("/", controller.AddPerson)
	group.Delete("/:id", controller.DeletePerson)
	// router.Get("/people", controller.ListPeople)
	// router.Post("/people", controller.AddPerson)
	// router.Delete("/people", middleware.New(middleware.Config{}), controller.DeletePerson)
}
/*
 ...
 */
```
**Note** that I
- originally authorized only one route, `/people/delete` and then ensured that all `/people` routes are authorized.

Lastly, I add logic to the that require authorization to ensure that the person role saved in the session matches the route required by the route.
``` go
/*
 ... people_controller.go / func (uc *PeopleController) DeletePerson(c *fiber.Ctx) error
 */
 	claims := c.Locals(constants.JWT_CLAIMS)

	if claims == nil {
		return c.SendString("Jwt was bypassed")
	}

	role := claims.(jwt.MapClaims)["role"]
	if role != "Admin" {
		// return c.Status(403).SendString("Only Admin users can delete people")
		middleware.Unauthorized(c)
		return c.Status(403).SendString("Unthorized: Only Adminstrators can delete people")
	}
/*
 ...
 */
```
**Note** that I
- use trhge 
  
# References
- [Golang Fiber Custom Middleware](https://strapengine.com/build-a-custom-jwt-go-fiber-middleware/) - A simple, but complete description;
- [Golang JWT](https://github.com/golang-jwt/jwt) - A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens;
- [Golang Crypto](https://pkg.go.dev/golang.org/x/crypto) - Password encryption / decryption 
- [My notes](https://share.evernote.com/note/bf91deec-398f-a660-2584-16da9095f3dd)