package domain

import "time"

// User represents a user entity in the domain
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

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"phone"`
	MembershipType string `json:"membership_type"`
	Points         int    `json:"points"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	Email          string `json:"email,omitempty" validate:"omitempty,email"`
	Phone          string `json:"phone,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
	Points         int    `json:"points,omitempty"`
}

// UserRepository defines the repository interface for user operations
type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

// UserUseCase defines the use case interface for user operations
type UserUseCase interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	CreateUser(req CreateUserRequest) (*User, error)
	UpdateUser(id uint, req UpdateUserRequest) (*User, error)
	DeleteUser(id uint) error
}
