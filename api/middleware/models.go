package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/models"
	"github.com/h00s-go/tiny-link-backend/services"
)

type ModelsMiddleware struct {
	Links *models.Links
}

func NewModelsMiddleware(services *services.Services) *ModelsMiddleware {
	return &ModelsMiddleware{
		Links: models.NewLinks(services),
	}
}

func (m *ModelsMiddleware) ModelsMiddleware(c *fiber.Ctx) error {
	c.Locals("models", m)
	return c.Next()
}

func GetModels(c *fiber.Ctx) *ModelsMiddleware {
	return c.Locals("models").(*ModelsMiddleware)
}
