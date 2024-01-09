package main

import (
	"log"

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

	// Database connection
	database, err := db.SetupDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Handlers
	handler := handlers.NewHandler(database)

	// Public Routes
	app.Get("/api/healthcheck", handler.HealthCheck)

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

	// Personalized Content Feed:
	app.Get("/api/secure/feed") // Fetch the personalized content feed based on user preferences and behavior.

	// Secured Routes

	//     User Feedback:
	//         POST /feedback - Submit feedback on content recommendations.

	//     Analytics (for User Engagement):
	//         GET /analytics/user-activity - Retrieve user activity and engagement analytics.

	//     Content Management (If you plan to allow user-generated content):
	//         POST /content/create - Create new content.
	//         PUT /content/update - Update existing content.
	//         DELETE /content/delete - Delete content.

	// Administrative Routes (If applicable)

	//     Admin Content Management:
	//         POST /admin/content - Add new content to the platform.
	//         PUT /admin/content/{id} - Update specific content.
	//         DELETE /admin/content/{id} - Remove specific content from the platform.

	//     Admin User Management:
	//         GET /admin/users - List all users.
	//         GET /admin/users/{id} - View details of a specific user.
	//         DELETE /admin/users/{id} - Delete a user account.

	//     Admin Analytics:
	//         GET /admin/analytics - Access overall platform analytics.

	// Start the scheduled fetching
	utils.StartScheduledFetching()

	// Start server
	log.Println("Server starting on port 4000...")
	log.Fatal(app.Listen(":4000"))
}

