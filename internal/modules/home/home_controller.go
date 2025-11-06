package home

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) GetIndex(c *fiber.Ctx) error {
	return c.Render("layouts/base", fiber.Map{
		"Title":      "Welcome to Go + Fiber + HTMX!",
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"PersonName": c.Locals("PersonName"),
		"Copyright":  "2025 Madrone Logistics",
	}, )
}
