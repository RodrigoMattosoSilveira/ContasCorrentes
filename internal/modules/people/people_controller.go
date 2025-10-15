package people

import (
	"log"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/validator"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type PeopleController struct {
	service *PeopleService
}

func NewPeopleController(service *PeopleService) *PeopleController {
	// forcing a change to test git
	return &PeopleController{service: service}
}

func (uc *PeopleController) ListPeople(c *fiber.Ctx) error {
	People, err := uc.service.GetAllPeople()
	if err != nil {
		return c.Status(500).SendString("Error fetching assopciates")
	}

	return c.Render("pages/people/index", fiber.Map{
		"Title":      "People Management",
		"People":  People,
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"PersonName":   c.Locals("PersonName"),
	}, "layouts/base")
}

func (uc *PeopleController) AddPerson(c *fiber.Ctx) error {
	log.Println("PeopleController - Adding Person")
	var person Person

	if err := c.BodyParser(&person); err != nil {
		log.Println("PeopleController - Body Parse Error: ", err)
		return c.Status(400).SendString("Invalid Person data")
	}

	if err := validator.ValidateStruct(person); err != nil {
		log.Println("PeopleController - Validation error ", err)
		return c.Status(400).SendString("Validation failed: " + err.Error())
	}

	hashedPassword, err := HashPassword(person.Password)
	if err != nil {
		return c.Status(400).SendString("Failed to hash password:")
	}
	person.Password = hashedPassword

	if err := uc.service.CreatePerson(&person); err != nil {
		log.Println("PeopleController - Unable to add record to database", err)
		return c.Status(500).SendString("Error creating Person")
	}

	return c.Render("partials/people/person-row", person)
}

func (uc *PeopleController) DeletePerson(c *fiber.Ctx) error {
	claimData := c.Locals("jwtClaims")

	if claimData == nil {
		return c.SendString("Jwt was bypassed")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	if err := uc.service.DeletePerson(uint(id)); err != nil {
		return c.Status(500).SendString("Error deleting Person")
	}

	return c.SendString("")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}