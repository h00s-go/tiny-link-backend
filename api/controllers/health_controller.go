package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
	services *Services
}

func NewHealthController(services *Services) *HealthController {
	return &HealthController{
		services: services,
	}
}

func (h *HealthController) GetHealthHandler(c echo.Context) error {
	h.services.Logger.Println("Health check")
	return c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
}
