package main

import (
	"log"

	"kbtg.tech/ai-backend-workshop/internal/config"
	"kbtg.tech/ai-backend-workshop/internal/handler"
	"kbtg.tech/ai-backend-workshop/internal/repository"
	"kbtg.tech/ai-backend-workshop/internal/usecase"
	"kbtg.tech/ai-backend-workshop/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize database
	db, err := database.NewDatabase(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Seed database
	if err := db.SeedData(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "*",
	}))

	// Setup routes
	setupRoutes(app, userHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}

func setupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	// API v1
	api := app.Group("/api/v1")

	// Health check endpoint
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "KBTG AI Backend Workshop is running!",
		})
	})

	// Hello World endpoint
	api.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}
		return c.JSON(fiber.Map{
			"message": "Hello, " + name + "!",
		})
	})

	// User routes
	users := api.Group("/users")
	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)

	// Static files
	app.Static("/", "./public")
}
