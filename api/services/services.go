package services

import (
	"log"

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
