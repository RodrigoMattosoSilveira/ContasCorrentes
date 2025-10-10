package people

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	service := NewPeopleService(db)
	controller := NewPeopleController(service)

	group := router.Group("/People")

	group.Get("/", controller.ListPeople)
	group.Post("/", controller.AddPerson)
	group.Delete("/:id", controller.DeletePerson)
}
