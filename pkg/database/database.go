package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kbtg.tech/ai-backend-workshop/internal/domain"
)

// DB holds the database connection
type DB struct {
	*gorm.DB
}

// NewDatabase creates a new database connection
func NewDatabase(dbPath string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate the User model
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &DB{db}, nil
}

// SeedData seeds the database with initial data
func (db *DB) SeedData() error {
	// Check if users already exist
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		return nil // Data already exists
	}

	seedUsers := []domain.User{
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
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to seed user: %w", err)
		}
	}

	log.Println("Database seeded with initial users")
	return nil
}

// GenerateMembershipID generates a random membership ID
func GenerateMembershipID() string {
	return fmt.Sprintf("LBK%06d", time.Now().UnixNano()%999999)
}
