package main

import (
	"log"
	"net/http"
	"os"

	"github.com/h00s-go/tiny-link-backend/api"
	"github.com/h00s-go/tiny-link-backend/config"
)

func main() {
	config := config.NewConfig()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	api := api.NewAPI(config, logger)

	logger.Println("Listening on :8080")
	http.ListenAndServe(":8080", api.NewRouter())
}
