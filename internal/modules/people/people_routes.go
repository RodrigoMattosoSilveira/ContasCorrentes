package people

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/middleware"
)

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

