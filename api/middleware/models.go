package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/models"
	"github.com/h00s-go/tiny-link-backend/api/services"
)

type Models struct {
	Links *models.Links
}

func NewModels(services *services.Services) *Models {
	return &Models{
		Links: models.NewLinks(services),
	}
}

func (m *Models) ModelsMiddleware(c *fiber.Ctx) error {
	c.Locals("models", m)
	return c.Next()
}

func GetModels(c *fiber.Ctx) *Models {
	return c.Locals("models").(*Models)
}
