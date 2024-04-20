package main

import (
	"fmt"
	"log"

	"github.com/LaMoldy/TheBowlShop/api"
	c "github.com/LaMoldy/TheBowlShop/config"
)

func main() {
	// Load the config for both the database and the server
	config, err := c.LoadConfig()
	if err != nil {
		log.Fatalf("Error: Failed to load config.\n%s", err)
	}

	// Create the router
	router, err := api.CreateApi()
	if err != nil {
		log.Fatalf("Error: Failed to create API.\n%s", err)
	}

	// Start listening
	serverConfig := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	router.Run(serverConfig)
}
