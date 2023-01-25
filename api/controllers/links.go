package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/models"
	"github.com/h00s-go/tiny-link-backend/api/services"
)

type LinksController struct {
	services *services.Services
}

func NewLinksController(services *services.Services) *LinksController {
	return &LinksController{
		services: services,
	}
}

func (l *LinksController) GetLinkByShortURIHandler(c *fiber.Ctx) error {
	link, err := models.GetLinkByShortURI(l.services, c.Params("short_uri"))
	if err != nil {
		return c.SendStatus(404)
	}
	return c.JSON(link)
}
