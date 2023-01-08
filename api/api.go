package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db"
	"github.com/labstack/echo/v4"
)

type API struct {
	config   *config.Config
	server   *echo.Echo
	services *Services
}

type Services struct {
	db     *db.Database
	logger *log.Logger
}

func NewAPI(config *config.Config, db *db.Database, logger *log.Logger) *API {
	return &API{
		config: config,
		server: echo.New(),
		services: &Services{
			db:     db,
			logger: logger,
		},
	}
}

func (api *API) Start() {
	api.services.logger.Println("Starting server on :8080")
	go func() {
		if err := api.server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			api.services.logger.Fatal("Error starting server: " + err.Error())
		}
	}()
}

func (api *API) WaitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := api.server.Shutdown(ctx); err != nil {
		api.services.logger.Fatal(err)
	}
}
