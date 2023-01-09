package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LinksController struct {
	services *Services
}

func NewLinksController(services *Services) *LinksController {
	return &LinksController{
		services: services,
	}
}

func (l *LinksController) GetLinkHandler(c echo.Context) error {
	l.services.Logger.Println("Get link")
	return c.JSON(http.StatusOK, map[string]string{
		"id":     c.Param("id"),
		"status": "OK",
	})
}
