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
	links := models.NewLinks(l.services)
	link, err := links.FindByShortURI(c.Params("short_uri"))
	if err != nil {
		return c.SendStatus(404)
	}
	return c.JSON(link)
}

func (l *LinksController) CreateLinkHandler(c *fiber.Ctx) error {
	links := models.NewLinks(l.services)
	link := &models.Link{}
	if err := c.BodyParser(link); err != nil {
		return c.SendStatus(400)
	}

	var err error
	if link, err = links.Create(link.URL); err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(link)
}
