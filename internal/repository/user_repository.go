package repository

import (
	"errors"

	"gorm.io/gorm"
	"kbtg.tech/ai-backend-workshop/internal/domain"
	"kbtg.tech/ai-backend-workshop/pkg/database"
)

// userRepository implements the UserRepository interface
type userRepository struct {
	db *database.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAll retrieves all users from the database
func (r *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// Create creates a new user in the database
func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user in the database
func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user by ID
func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
