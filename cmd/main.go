package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"Newsly/config"
	"Newsly/internal/app"
	"Newsly/internal/db"
	//"Newsly/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config := config.MustLoadConfig()

	// Database connection
	database := db.MustLoadDatabase(config.DBConnectionString)

	// Start the scheduled fetching
	// go utils.StartScheduledFetching(database, config.NewsAPIKey)
	// go utils.StoreNews(database)

	// Start the server
	app := app.MustStart(database, config)

	// Handle graceful shutdown
	handleGracefulShutdown(app)

}

// handleGracefulShutdown waits for interrupt signal and shuts down the server
func handleGracefulShutdown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-c
	log.Println("Gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
		panic(err)
	}
	log.Println("Server stopped")
}