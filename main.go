package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database connection
var db *gorm.DB

// User model based on the profile UI
type User struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	FirstName      string    `json:"first_name" gorm:"not null"`
	LastName       string    `json:"last_name" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Phone          string    `json:"phone"`
	MembershipType string    `json:"membership_type" gorm:"default:'Bronze'"` // Bronze, Silver, Gold
	MembershipID   string    `json:"membership_id" gorm:"unique"`
	JoinDate       time.Time `json:"join_date" gorm:"autoCreateTime"`
	Points         int       `json:"points" gorm:"default:0"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Generate a random membership ID
func generateMembershipID() string {
	return fmt.Sprintf("LBK%06d", rand.Intn(999999))
}

func main() {
	// Initialize database
	initDatabase()

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

func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the User model
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed initial data if no users exist
	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		seedUsers := []User{
			{
				FirstName:      "สมชาย",
				LastName:       "ใจดี",
				Email:          "somchai@example.com",
				Phone:          "081-234-5678",
				MembershipType: "Gold",
				MembershipID:   "LBK001234",
				JoinDate:       time.Now().AddDate(-1, 0, 0), // 1 year ago
				Points:         15420,
			},
			{
				FirstName:      "สมหญิง",
				LastName:       "รักดี",
				Email:          "somying@example.com",
				Phone:          "089-765-4321",
				MembershipType: "Silver",
				MembershipID:   "LBK001235",
				JoinDate:       time.Now().AddDate(0, -6, 0), // 6 months ago
				Points:         8750,
			},
		}

		for _, user := range seedUsers {
			db.Create(&user)
		}

		log.Println("Database seeded with initial users")
	}
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

func getUsers(c *fiber.Ctx) error {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"count": len(users),
	})
}

func getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var user User
	result := db.First(&user, uint(id))
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func createUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "First name, last name, and email are required",
		})
	}

	// Generate membership ID if not provided
	if user.MembershipID == "" {
		user.MembershipID = generateMembershipID()
	}

	// Set default membership type if not provided
	if user.MembershipType == "" {
		user.MembershipType = "Bronze"
	}

	result := db.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

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

	var user User
	result := db.First(&user, uint(id))
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})
	}

	var updatedData User
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update fields
	if updatedData.FirstName != "" {
		user.FirstName = updatedData.FirstName
	}
	if updatedData.LastName != "" {
		user.LastName = updatedData.LastName
	}
	if updatedData.Email != "" {
		user.Email = updatedData.Email
	}
	if updatedData.Phone != "" {
		user.Phone = updatedData.Phone
	}
	if updatedData.MembershipType != "" {
		user.MembershipType = updatedData.MembershipType
	}
	if updatedData.Points != 0 {
		user.Points = updatedData.Points
	}

	result = db.Save(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func deleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	result := db.Delete(&User{}, uint(id))
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
