package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/tiny-link-backend/api/services"
	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db"
)

type API struct {
	config   *config.Config
	server   *fiber.App
	services *services.Services
}

func NewAPI(config *config.Config, database *db.Database, memStore *db.MemStore, logger *log.Logger) *API {
	return &API{
		config: config,
		server: fiber.New(),
		services: services.NewServices(
			database,
			memStore,
			logger,
		),
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
