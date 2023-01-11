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
