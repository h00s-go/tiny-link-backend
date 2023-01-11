package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/services"
)

type HealthController struct {
	services *services.Services
}

func NewHealthController(services *services.Services) *HealthController {
	return &HealthController{
		services: services,
	}
}

func (h *HealthController) GetHealthHandler(c *fiber.Ctx) error {
	h.services.Logger.Println("Health check")
	return c.JSON(fiber.Map{
		"status": "OK",
		"age":    20,
	})
}
