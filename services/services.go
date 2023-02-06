package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/db"
)

type Services struct {
	DB     *db.Database
	IMDS   *db.MemStore
	Logger *log.Logger
}

func NewServices(db *db.Database, imds *db.MemStore, logger *log.Logger) *Services {
	return &Services{
		DB:     db,
		IMDS:   imds,
		Logger: logger,
	}
}

func (s *Services) ServicesMiddleware(c *fiber.Ctx) error {
	c.Locals("services", s)
	return c.Next()
}

func GetServices(c *fiber.Ctx) *Services {
	return c.Locals("services").(*Services)
}
