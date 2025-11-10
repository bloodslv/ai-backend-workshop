package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Create Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "KBTG AI Backend Workshop",
	})

	// Add middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "*",
	}))

	// Routes
	setupRoutes(app)

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func setupRoutes(app *fiber.App) {
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
	users.Get("/", getUsers)
	users.Get("/:id", getUser)
	users.Post("/", createUser)
	users.Put("/:id", updateUser)
	users.Delete("/:id", deleteUser)

	// Static files
	app.Static("/", "./public")
}

// In-memory storage for demo (replace with database in production)
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}

func getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	for _, user := range users {
		if user.ID == id {
			return c.JSON(fiber.Map{
				"data": user,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "User not found",
	})
}

func createUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Generate new ID
	user.ID = len(users) + 1
	users = append(users, user)

	return c.Status(201).JSON(fiber.Map{
		"data": user,
	})
}

func updateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			return c.JSON(fiber.Map{
				"data": updatedUser,
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "User not found",
	})
}

func deleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(fiber.Map{
				"message": "User deleted successfully",
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "User not found",
	})
}
