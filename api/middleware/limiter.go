package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/h00s-go/tiny-link-backend/config"
)

type LimiterMiddleware struct {
	config *config.Limiter
}

func NewLimiterMiddleware(config *config.Limiter) *LimiterMiddleware {
	return &LimiterMiddleware{
		config: config,
	}
}

func (l *LimiterMiddleware) LimiterMiddleware() func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:          l.config.Max,
		Expiration:   time.Duration(l.config.Expiration) * time.Second,
		Next:         l.throttling,
		LimitReached: l.throttleClient,
	})
}

func (l *LimiterMiddleware) throttling(c *fiber.Ctx) bool {
	return c.Method() != "POST"
}

func (l *LimiterMiddleware) throttleClient(c *fiber.Ctx) error {
	GetServices(c).Logger.Println("Throttling client: " + c.IP())
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"error": "Too many requests",
	})
}
