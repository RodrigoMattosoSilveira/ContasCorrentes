package associates

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	service := NewAssociatesService(db)
	controller := NewAssociatesController(service)

	group := router.Group("/associates")

	group.Get("/", controller.ListAssociates)
	group.Post("/", controller.AddAssociate)
	group.Delete("/:id", controller.DeleteAssociate)
}
