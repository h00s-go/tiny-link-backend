package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/services"
)

type ServicesMiddleware struct {
	Services *services.Services
}

func NewServicesMiddleware(services *services.Services) *ServicesMiddleware {
	return &ServicesMiddleware{
		Services: services,
	}
}

func (sm *ServicesMiddleware) ServicesMiddleware(c *fiber.Ctx) error {
	c.Locals("services", sm.Services)
	return c.Next()
}
