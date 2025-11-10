package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kbtg.tech/ai-backend-workshop/internal/domain"
	"kbtg.tech/ai-backend-workshop/internal/mocks"
)

func TestUserUseCase_GetAllUsers(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	expectedUsers := []domain.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
	}

	mockRepo.On("GetAll").Return(expectedUsers, nil)

	// Act
	result, err := useCase.GetAllUsers()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetAllUsers_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	mockRepo.On("GetAll").Return([]domain.User{}, errors.New("database error"))

	// Act
	result, err := useCase.GetAllUsers()

	// Assert
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetUserByID(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	expectedUser := &domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	mockRepo.On("GetByID", uint(1)).Return(expectedUser, nil)

	// Act
	result, err := useCase.GetUserByID(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetUserByID_InvalidID(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	// Act
	result, err := useCase.GetUserByID(0)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "invalid user ID", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	req := domain.CreateUserRequest{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		Points:         100,
	}

	mockRepo.On("GetByEmail", "john@example.com").Return(nil, errors.New("user not found"))
	mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

	// Act
	result, err := useCase.CreateUser(req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.FirstName, result.FirstName)
	assert.Equal(t, req.LastName, result.LastName)
	assert.Equal(t, req.Email, result.Email)
	assert.Equal(t, req.Phone, result.Phone)
	assert.Equal(t, req.MembershipType, result.MembershipType)
	assert.Equal(t, req.Points, result.Points)
	assert.NotEmpty(t, result.MembershipID)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_CreateUser_MissingRequiredFields(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	req := domain.CreateUserRequest{
		FirstName: "John",
		// LastName missing
		Email: "john@example.com",
	}

	// Act
	result, err := useCase.CreateUser(req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "first name, last name, and email are required", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_CreateUser_EmailExists(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	req := domain.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	existingUser := &domain.User{ID: 1, Email: "john@example.com"}
	mockRepo.On("GetByEmail", "john@example.com").Return(existingUser, nil)

	// Act
	result, err := useCase.CreateUser(req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "user with this email already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_UpdateUser(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	existingUser := &domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Points:    100,
	}

	updateReq := domain.UpdateUserRequest{
		FirstName: "Jane",
		Points:    200,
	}

	mockRepo.On("GetByID", uint(1)).Return(existingUser, nil)
	mockRepo.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)

	// Act
	result, err := useCase.UpdateUser(1, updateReq)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)           // unchanged
	assert.Equal(t, "john@example.com", result.Email) // unchanged
	assert.Equal(t, 200, result.Points)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_DeleteUser(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	existingUser := &domain.User{ID: 1, Email: "john@example.com"}
	mockRepo.On("GetByID", uint(1)).Return(existingUser, nil)
	mockRepo.On("Delete", uint(1)).Return(nil)

	// Act
	err := useCase.DeleteUser(1)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_DeleteUser_UserNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	mockRepo.On("GetByID", uint(1)).Return(nil, errors.New("user not found"))

	// Act
	err := useCase.DeleteUser(1)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}
