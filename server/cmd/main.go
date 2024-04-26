package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Anisulh/content_personalization/db"
	"github.com/Anisulh/content_personalization/handlers"
	"github.com/Anisulh/content_personalization/middleware"
	"github.com/Anisulh/content_personalization/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load environment variables
	utils.LoadEnv()

	// Database connection
	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	
	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: utils.ClientAddress,
		AllowHeaders: "Authorization, Origin, Content-Type, Accept",
	})) 
	app.Use(logger.New()) 
	app.Use("/api/secure", middleware.JWTProtected())

	
	// Start the scheduled fetching
	go utils.StartScheduledFetching(database)
	go utils.StoreNews(database)

	// Handlers
	handler := handlers.NewHandler(database)

	

	// Public Routes
	app.Get("/api/health-check", handler.HealthCheck)

	// User Auth
	app.Post("/api/register", handler.UserRegistration)
	app.Post("/api/login", handler.UserLogin)

	// Content Discover
	app.Get("/api/content", handler.GetContent)
	app.Get("/api/content/categories", handler.GetContentCategories)


	// Secured Routes

	// User Profile
	app.Get("/api/secure/user/profile", handler.GetUserProfile)
	app.Put("/api/secure/user/profile", handler.UpdateUserProfile)
	app.Get("/api/secure/user/preferences", handler.GetUserPreferences)
	app.Put("/api/secure/user/preferences", handler.UpdateUserPreferences)

	// Content Interaction
	app.Post("/api/secure/content/:contentId/like", handlers.LikeContent) 
	app.Post("/api/secure/content/:contentId/dislike", handlers.DislikeContent)  
	app.Post("/api/secure/content/:contentId/bookmark", handlers.BookmarkContent)  


	// Channel to listen for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Println("Server starting on port 4000...")
		if err := app.Listen(":4000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Block until a signal is received
	<-c
	log.Println("Gracefully shutting down...")


	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	// Additional resource cleanup here...
	log.Println("Server stopped")

}

