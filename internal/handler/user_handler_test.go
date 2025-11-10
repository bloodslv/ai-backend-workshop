package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"kbtg.tech/ai-backend-workshop/internal/domain"
	"kbtg.tech/ai-backend-workshop/internal/mocks"
)

func setupTestApp() *fiber.App {
	return fiber.New()
}

func TestUserHandler_GetUsers(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	expectedUsers := []domain.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
	}

	mockUseCase.On("GetAllUsers").Return(expectedUsers, nil)

	app.Get("/users", handler.GetUsers)

	// Act
	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_GetUsers_Error(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	mockUseCase.On("GetAllUsers").Return([]domain.User{}, errors.New("database error"))

	app.Get("/users", handler.GetUsers)

	// Act
	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_GetUser(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	expectedUser := &domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	mockUseCase.On("GetUserByID", uint(1)).Return(expectedUser, nil)

	app.Get("/users/:id", handler.GetUser)

	// Act
	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_GetUser_InvalidID(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	app.Get("/users/:id", handler.GetUser)

	// Act
	req := httptest.NewRequest("GET", "/users/invalid", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_GetUser_NotFound(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	mockUseCase.On("GetUserByID", uint(1)).Return(nil, errors.New("user not found"))

	app.Get("/users/:id", handler.GetUser)

	// Act
	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_CreateUser(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	createReq := domain.CreateUserRequest{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		Points:         100,
	}

	expectedUser := &domain.User{
		ID:             1,
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		Points:         100,
	}

	mockUseCase.On("CreateUser", createReq).Return(expectedUser, nil)

	app.Post("/users", handler.CreateUser)

	// Act
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_CreateUser_InvalidBody(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	app.Post("/users", handler.CreateUser)

	// Act
	req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_CreateUser_ValidationError(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	createReq := domain.CreateUserRequest{
		FirstName: "John",
		// LastName missing
		Email: "john@example.com",
	}

	mockUseCase.On("CreateUser", createReq).Return(nil, errors.New("first name, last name, and email are required"))

	app.Post("/users", handler.CreateUser)

	// Act
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	updateReq := domain.UpdateUserRequest{
		FirstName: "Jane",
		Points:    200,
	}

	expectedUser := &domain.User{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "john@example.com",
		Points:    200,
	}

	mockUseCase.On("UpdateUser", uint(1), updateReq).Return(expectedUser, nil)

	app.Put("/users/:id", handler.UpdateUser)

	// Act
	body, _ := json.Marshal(updateReq)
	req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_DeleteUser(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	mockUseCase.On("DeleteUser", uint(1)).Return(nil)

	app.Delete("/users/:id", handler.DeleteUser)

	// Act
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}

func TestUserHandler_DeleteUser_NotFound(t *testing.T) {
	// Arrange
	mockUseCase := new(mocks.MockUserUseCase)
	handler := NewUserHandler(mockUseCase)
	app := setupTestApp()

	mockUseCase.On("DeleteUser", uint(1)).Return(errors.New("user not found"))

	app.Delete("/users/:id", handler.DeleteUser)

	// Act
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	mockUseCase.AssertExpectations(t)
}
