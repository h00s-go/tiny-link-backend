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
	config *config.Config
	db     *db.Database
	logger *log.Logger
	server *echo.Echo
}

func NewAPI(config *config.Config, db *db.Database, logger *log.Logger) *API {
	return &API{
		config: config,
		db:     db,
		logger: logger,
		server: echo.New(),
	}
}

func (api *API) Start() {
	api.logger.Println("Starting server on :8080")
	go func() {
		if err := api.server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			api.logger.Fatal("Error starting server: " + err.Error())
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
		api.logger.Fatal(err)
	}
}
