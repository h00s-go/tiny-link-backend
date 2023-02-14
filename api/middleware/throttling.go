package middleware

import "github.com/gofiber/fiber/v2"

func Throttling(c *fiber.Ctx) bool {
	return c.Method() != "POST"
}
