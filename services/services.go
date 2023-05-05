package services

import (
	"log"

	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db"
)

type Services struct {
	DB     *db.Database
	IMDS   *db.MemStore
	Logger *log.Logger
}

func NewServices(config *config.Config, logger *log.Logger) *Services {
	database := db.NewDatabase(&config.Database)
	if err := database.Connect(); err != nil {
		logger.Fatal(err)
	}
	if err := database.Migrate(); err != nil {
		logger.Fatal(err)
	}

	memStore := db.NewMemStore(&config.MemStore)
	if err := memStore.Connect(); err != nil {
		logger.Fatal(err)
	}

	return &Services{
		DB:     database,
		IMDS:   memStore,
		Logger: logger,
	}
}

func (s *Services) Close() {
	s.DB.Close()
	s.IMDS.Close()
}
