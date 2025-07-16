package main

import (
	"foods-drinks-app/config"
	"foods-drinks-app/models"
	"log"
)

func main() {
	config.ConnectDatabase()

	if err := models.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Your router setup, routes, etc.
}
