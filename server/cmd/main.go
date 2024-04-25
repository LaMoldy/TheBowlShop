package main

import (
	"fmt"
	"log"

	"github.com/LaMoldy/TheBowlShop/api"
	c "github.com/LaMoldy/TheBowlShop/config"
	db "github.com/LaMoldy/TheBowlShop/internal/database"
)

func main() {
	// Load the config for both the database and the server
	config, err := c.LoadConfig()
	if err != nil {
		log.Fatalf("Error: Failed to load config.\n%s", err)
	}

	// Connect to the database
	database, err := db.CreateDatabaseConnection(config)
	if err != nil {
		log.Fatalf("Error: Failed to create database connection.\n%s", err)
	}

	// Close the database when it's done being used
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Error: Failed to close database connection\n%s", err)
	}
	defer sqlDB.Close()

	// Create the router
	router, err := api.CreateApi(database)
	if err != nil {
		log.Fatalf("Error: Failed to create API.\n%s", err)
	}

	// Start listening
	serverConfig := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	router.Run(serverConfig)
}
