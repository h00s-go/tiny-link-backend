package controllers

import (
	"github.com/h00s-go/tiny-link-backend/api/models"
	"github.com/h00s-go/tiny-link-backend/api/services"
	"github.com/labstack/echo/v4"
)

type LinksController struct {
	services *services.Services
}

func NewLinksController(services *services.Services) *LinksController {
	return &LinksController{
		services: services,
	}
}

func (l *LinksController) GetLinkHandler(c echo.Context) error {
	link, err := models.GetLinkByID(l.services, c.Param("id"))
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, link)
}
