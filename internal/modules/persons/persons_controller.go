package persons

import (
	"log"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type PersonsController struct {
	service *PersonsService
}

func NewPersonsController(service *PersonsService) *PersonsController {
	// forcing a change to test git
	return &PersonsController{service: service}
}

func (uc *PersonsController) ListPersons(c *fiber.Ctx) error {
	Persons, err := uc.service.GetAllPersons()
	if err != nil {
		return c.Status(500).SendString("Error fetching assopciates")
	}

	return c.Render("pages/persons/index", fiber.Map{
		"Title":      "Persons Management",
		"Persons":  Persons,
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"PersonName":   c.Locals("PersonName"),
	}, "layouts/base")
}

func (uc *PersonsController) AddPerson(c *fiber.Ctx) error {
	var Person Person
	log.Println("Adding Person")
	if err := c.BodyParser(&Person); err != nil {
		return c.Status(400).SendString("Invalid Person data")
	}

	if err := validator.ValidateStruct(Person); err != nil {
		return c.Status(400).SendString("Validation failed: " + err.Error())
	}

	if err := uc.service.CreatePerson(&Person); err != nil {
		return c.Status(500).SendString("Error creating Person")
	}

	return c.Render("partials/Persons/Person-row", Person)
}

func (uc *PersonsController) DeletePerson(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	if err := uc.service.DeletePerson(uint(id)); err != nil {
		return c.Status(500).SendString("Error deleting Person")
	}

	return c.SendString("")
}