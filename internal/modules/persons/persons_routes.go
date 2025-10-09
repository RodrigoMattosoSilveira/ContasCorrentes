package persons

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	service := NewPersonsService(db)
	controller := NewPersonsController(service)

	group := router.Group("/persons")

	group.Get("/", controller.ListPersons)
	group.Post("/", controller.AddPerson)
	group.Delete("/:id", controller.DeletePerson)
}
