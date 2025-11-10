package usecase

import (
	"errors"

	"kbtg.tech/ai-backend-workshop/internal/domain"
	"kbtg.tech/ai-backend-workshop/pkg/database"
)

// userUseCase implements the UserUseCase interface
type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// GetAllUsers retrieves all users
func (u *userUseCase) GetAllUsers() ([]domain.User, error) {
	return u.userRepo.GetAll()
}

// GetUserByID retrieves a user by ID
func (u *userUseCase) GetUserByID(id uint) (*domain.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user ID")
	}
	return u.userRepo.GetByID(id)
}

// CreateUser creates a new user
func (u *userUseCase) CreateUser(req domain.CreateUserRequest) (*domain.User, error) {
	// Validate required fields
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		return nil, errors.New("first name, last name, and email are required")
	}

	// Check if user with email already exists
	existingUser, _ := u.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create new user
	user := &domain.User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		MembershipType: req.MembershipType,
		Points:         req.Points,
		MembershipID:   database.GenerateMembershipID(),
	}

	// Set default membership type if not provided
	if user.MembershipType == "" {
		user.MembershipType = "Bronze"
	}

	err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user
func (u *userUseCase) UpdateUser(id uint, req domain.UpdateUserRequest) (*domain.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user ID")
	}

	// Get existing user
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if email is being changed to an existing email
	if req.Email != "" && req.Email != user.Email {
		existingUser, _ := u.userRepo.GetByEmail(req.Email)
		if existingUser != nil {
			return nil, errors.New("user with this email already exists")
		}
		user.Email = req.Email
	}

	// Update fields if provided
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.MembershipType != "" {
		user.MembershipType = req.MembershipType
	}
	if req.Points != 0 {
		user.Points = req.Points
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (u *userUseCase) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid user ID")
	}

	// Check if user exists
	_, err := u.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return u.userRepo.Delete(id)
}
