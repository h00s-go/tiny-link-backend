package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/middleware"
	"github.com/h00s-go/tiny-link-backend/api/models"
)

func GetLinkByShortURIHandler(c *fiber.Ctx) error {
	links := middleware.GetModels(c).Links
	link, err := links.FindByShortURI(c.Params("short_uri"))
	if err != nil {
		return c.SendStatus(404)
	}
	return c.JSON(link)
}

func RedirectLinkByShortURIHandler(c *fiber.Ctx) error {
	links := middleware.GetModels(c).Links
	link, err := links.FindByShortURI(c.Params("short_uri"))
	if err != nil {
		return c.SendStatus(404)
	}
	return c.Redirect(link.URL, 301)
}

func CreateLinkHandler(c *fiber.Ctx) error {
	links := middleware.GetModels(c).Links
	link := new(models.Link)
	if err := c.BodyParser(link); err != nil {
		return c.SendStatus(400)
	}

	var err error
	if link, err = links.Create(link.URL); err != nil {
		middleware.GetServices(c).Logger.Println("Error creating link: ", err)
		return c.SendStatus(500)
	} else {
		c.JSON(link)
		return c.SendStatus(201)
	}
}
