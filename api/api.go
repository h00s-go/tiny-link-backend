package api

import (
	"log"

	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db"
)

type API struct {
	config *config.Config
	db     *db.Database
	logger *log.Logger
}

func NewAPI(config *config.Config, db *db.Database, logger *log.Logger) *API {
	return &API{
		config: config,
		db:     db,
		logger: logger,
	}
}
