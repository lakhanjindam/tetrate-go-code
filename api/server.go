package api

import (
	"fmt"
	"log"
)

// Run starts the API server
func Run(portNumber string) error {
	r := NewRouter()

	address := fmt.Sprintf(":%s", portNumber)
	log.Println("Running Gin REST API on address:", address)

	//! Starts the api server(blocking call for main goroutine)
	err := r.Run(address)
	if err != nil {
		return err
	}

	return nil
}
