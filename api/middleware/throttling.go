package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Throttling(c *fiber.Ctx) bool {
	return c.Method() != "POST"
}

func ThrottleClient(c *fiber.Ctx) error {
	GetServices(c).Logger.Println("Throttling client: " + c.IP())
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"error": "Too many requests",
	})
}
