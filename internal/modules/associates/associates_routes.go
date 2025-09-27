package associates

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	service := NewAssociatesService(db)
	controller := NewAssociatesController(service)

	usersGroup := router.Group("/associates")

	usersGroup.Get("/", controller.ListAssociates)
	usersGroup.Post("/", controller.addAssociate)
	usersGroup.Delete("/:id", controller.DeleteAssociate)
}
