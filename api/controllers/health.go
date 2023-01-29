package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/services"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) GetHealthHandler(c *fiber.Ctx) error {
	s := services.GetServices(c)
	s.Logger.Println("Health check")
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
