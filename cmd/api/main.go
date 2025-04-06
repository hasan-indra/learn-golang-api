package main

import (
	"fmt"
	"github.com/hasan-indra/learn-golang-api/internal/config"
	"github.com/hasan-indra/learn-golang-api/internal/routes"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin router
	router := routes.SetupRouter()

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server running on %s", serverAddr)
	router.Run(serverAddr)
}
