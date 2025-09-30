package associates

import (
	"log"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type AssociatesController struct {
	service *AssociatesService
}

func NewAssociatesController(service *AssociatesService) *AssociatesController {
	return &AssociatesController{service: service}
}

func (uc *AssociatesController) ListAssociates(c *fiber.Ctx) error {
	associates, err := uc.service.GetAllAssociates()
	if err != nil {
		return c.Status(500).SendString("Error fetching assopciates")
	}

	return c.Render("pages/associates/index", fiber.Map{
		"Title":      "Associates Management",
		"Associates":  associates,
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Associatename":   c.Locals("Associatename"),
	}, "layouts/base")
}

func (uc *AssociatesController) AddAssociate(c *fiber.Ctx) error {
	var associate Associate
	log.Println("Adding associate")
	if err := c.BodyParser(&associate); err != nil {
		return c.Status(400).SendString("Invalid associate data")
	}

	if err := validator.ValidateStruct(associate); err != nil {
		return c.Status(400).SendString("Validation failed: " + err.Error())
	}

	if err := uc.service.CreateAssociate(&associate); err != nil {
		return c.Status(500).SendString("Error creating associate")
	}

	return c.Render("partials/associates/associate-row", associate)
}

func (uc *AssociatesController) DeleteAssociate(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	if err := uc.service.DeleteAssociate(uint(id)); err != nil {
		return c.Status(500).SendString("Error deleting associate")
	}

	return c.SendString("")
}
