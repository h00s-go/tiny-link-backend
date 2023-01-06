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

	db := db.NewDatabase(&config.Database)
	if err := db.Connect(); err != nil {
		logger.Fatal(err)
	}
	if err := db.Migrate(); err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	api := api.NewAPI(config, db, logger)
	api.SetRoutes()
	api.Start()
	api.WaitForShutdown()
}
