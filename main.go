package main

import (
	"log"

	"github.com/tetrate-go-code/api"
	"github.com/tetrate-go-code/config"
)

func main() {

	// Load the configuration
	if err := config.LoadConfig("config"); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	log.Println("Config:", config.Config)

	// Start the API server
	if err := api.Run(config.Config.Port); err != nil {
		log.Fatalf("error running api server: %v", err)
	}
}
