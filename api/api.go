package api

import (
	"log"

	"github.com/h00s-go/tiny-link-backend/config"
)

type API struct {
	config *config.Config
	logger *log.Logger
}

func NewAPI(config *config.Config, logger *log.Logger) *API {
	return &API{
		config: config,
		logger: logger,
	}
}
