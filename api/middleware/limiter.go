package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func LimiterMiddleware() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:          10,
		Expiration:   time.Minute,
		Next:         throttling,
		LimitReached: throttleClient,
	})
}

func throttling(c *fiber.Ctx) bool {
	return c.Method() != "POST"
}

func throttleClient(c *fiber.Ctx) error {
	GetServices(c).Logger.Println("Throttling client: " + c.IP())
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"error": "Too many requests",
	})
}
