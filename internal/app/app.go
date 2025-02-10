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
	// Handlers
	handler := handlers.NewHandler(db, config.JWTSecret, config.Environment, config.NewsAPIKey)

	// Middleware
	middleware := middleware.NewMiddleware(db, config.JWTSecret, config.Environment)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Serve static files
	app.Static("/static", "./web/static")

	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.ClientAddress,
		AllowHeaders: "Authorization, Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	// health check
	app.Get("/api/health-check", handler.HealthCheck)

	// Public Page Routes
	pageRouter := app.Group("/")
	pageRouter.Get("/", handler.GetHomePage)
	pageRouter.Get("/login", handler.GetLoginPage)
	pageRouter.Get("/register", handler.GetRegisterPage)
	

	// Private Page Routes
	privatePageRouter := app.Group("/auth", middleware.JWTProtected())
	privatePageRouter.Get("/interest-topics", handler.GetInterestsPage)
	privatePageRouter.Get("/feed", handler.GetFeedPage)

	// API Routes
	// User Auth
	userPublicRouter := app.Group("/api/v1/user")
	userPublicRouter.Post("/register", handler.UserRegistration)
	userPublicRouter.Post("/login", handler.UserLogin)
	userPublicRouter.Post("/logout", handler.UserLogout)

	// Content Discover
	// contentPublicRouter := app.Group("/api/content")
	// contentPublicRouter.Get("/", handler.GetContent)
	// contentPublicRouter.Get("/categories", handler.GetContentCategories)

	// Secured Routes
	// User Profile
	userPrivateRouter := app.Group("/api/v1/secure/user", middleware.JWTProtected())
	userPrivateRouter.Post("/interest-topics", handler.SaveUserInterests)

	return app, nil
}

func MustStart(db *gorm.DB, config *config.Config) *fiber.App {
	// Start server in a goroutine
	app, err := newApp(db, config)
	if err != nil {
		panic(err)
	}
	go func() {
		log.Println("Server starting on port:", config.Port)
		if err := app.Listen(":" + config.Port); err != nil {
			panic("Error starting server: " + err.Error())
		}
	}()
	return app
}
