package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/models"
	"github.com/h00s-go/tiny-link-backend/api/services"
)

func GetLinkByShortURIHandler(c *fiber.Ctx) error {
	links := models.NewLinks(services.GetServices(c))
	link, err := links.FindByShortURI(c.Params("short_uri"))
	if err != nil {
		return c.SendStatus(404)
	}
	return c.JSON(link)
}

func CreateLinkHandler(c *fiber.Ctx) error {
	s := services.GetServices(c)
	links := models.NewLinks(s)
	link := new(models.Link)
	if err := c.BodyParser(link); err != nil {
		return c.SendStatus(400)
	}

	var err error
	if link, err = links.Create(link.URL); err != nil {
		s.Logger.Println("Error creating link: ", err)
		return c.SendStatus(500)
	} else {
		c.JSON(link)
		return c.SendStatus(201)
	}
}
