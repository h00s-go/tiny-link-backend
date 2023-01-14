package main

import (
	"log"
	"os"

	"github.com/h00s-go/tiny-link-backend/api"
	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db"
)

func main() {
	config := config.NewConfig()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	database := db.NewDatabase(&config.Database)
	if err := database.Connect(); err != nil {
		logger.Fatal(err)
	}
	if err := database.Migrate(); err != nil {
		logger.Fatal(err)
	}
	defer database.Close()

	memstore := db.NewMemStore(&config.MemStore)
	if err := memstore.Connect(); err != nil {
		logger.Fatal(err)
	}
	defer memstore.Close()

	api := api.NewAPI(config, database, memstore, logger)
	api.Start()
	api.WaitForShutdown()
}
