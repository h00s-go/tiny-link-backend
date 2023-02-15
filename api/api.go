package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/h00s-go/tiny-link-backend/api/middleware"
	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db"
	"github.com/h00s-go/tiny-link-backend/services"
)

type API struct {
	config   *config.Config
	server   *fiber.App
	services *services.Services
}

func NewAPI(config *config.Config, database *db.Database, memStore *db.MemStore, logger *log.Logger) *API {
	server := fiber.New()
	services := services.NewServices(database, memStore, logger)
	servicesMiddleware := middleware.NewServicesMiddleware(services)
	modelsMiddleware := middleware.NewModelsMiddleware(services)

	server.Use(servicesMiddleware.ServicesMiddleware)
	server.Use(modelsMiddleware.ModelsMiddleware)
	server.Use(limiter.New(limiter.Config{
		Next:       middleware.Throttling,
		Max:        10,
		Expiration: time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			services.Logger.Println("Throttling client: " + c.IP())
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))

	return &API{
		config:   config,
		server:   server,
		services: services,
	}
}

func (api *API) Start() {
	api.services.Logger.Println("Starting server on :8080")
	api.setRoutes()
	go func() {
		if err := api.server.Listen(":8080"); err != nil && err != http.ErrServerClosed {
			api.services.Logger.Fatal("Error starting server: " + err.Error())
		}
	}()
}

func (api *API) WaitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := api.server.Shutdown(); err != nil {
		api.services.Logger.Fatal(err)
	}
}
