package app

import (
	"Newsly/config"
	"Newsly/internal/handlers"
	"Newsly/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)


func newApp(db *gorm.DB, config *config.Config) (*fiber.App, error) {

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.ClientAddress,
		AllowHeaders: "Authorization, Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())
	app.Use("/api/secure", middleware.JWTProtected())

	// Handlers
	handler := handlers.NewHandler(db)

	// Public Routes
	app.Get("/api/health-check", handler.HealthCheck)

	// User Auth
	userPublicRouter := app.Group("/api/user")
	userPublicRouter.Post("/register", handler.UserRegistration)
	userPublicRouter.Post("/login", handler.UserLogin)

	// Content Discover
	contentPublicRouter := app.Group("/api/content")
	contentPublicRouter.Get("/", handler.GetContent)
	contentPublicRouter.Get("/categories", handler.GetContentCategories)

	// Secured Routes

	// User Profile
	userPrivateRouter := app.Group("/api/user", middleware.JWTProtected())
	userPrivateRouter.Get("/profile", handler.GetUserProfile)
	userPrivateRouter.Put("/profile", handler.UpdateUserProfile)
	userPrivateRouter.Get("/preferences", handler.GetUserPreferences)
	userPrivateRouter.Put("/preferences", handler.UpdateUserPreferences)

	// Content Interaction
	contentPrivateRouter := app.Group("/api/content")
	contentPrivateRouter.Post("/:contentId/like", handlers.LikeContent)
	contentPrivateRouter.Post("/:contentId/dislike", handlers.DislikeContent)
	contentPrivateRouter.Post("/:contentId/bookmark", handlers.BookmarkContent)

	return app, nil
}

func MustStart(db *gorm.DB, config *config.Config) *fiber.App {
	// Start server in a goroutine
	app, err := newApp(db, config)
	if err != nil {
		panic(err)
	}
	go func() {
		log.Println("Server starting on port 4000...")
		if err := app.Listen(":" + config.Port); err != nil {
			panic("Error starting server: " + err.Error())
		}
	}()
	return app
}
