package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/services"
)

type Models struct {
	Links *Links
}

func NewModels(services *services.Services) *Models {
	return &Models{
		Links: NewLinks(services),
	}
}

func (m *Models) ModelsMiddleware(c *fiber.Ctx) error {
	c.Locals("models", m)
	return c.Next()
}

func GetModels(c *fiber.Ctx) *Models {
	return c.Locals("models").(*Models)
}
