package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"kbtg.tech/ai-backend-workshop/internal/config"
	"kbtg.tech/ai-backend-workshop/internal/domain"
	"kbtg.tech/ai-backend-workshop/internal/handler"
	"kbtg.tech/ai-backend-workshop/internal/repository"
	"kbtg.tech/ai-backend-workshop/internal/usecase"
	"kbtg.tech/ai-backend-workshop/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type APITestSuite struct {
	suite.Suite
	app    *fiber.App
	db     *database.DB
	config *config.Config
}

func (suite *APITestSuite) SetupSuite() {
	// Create test database
	gormDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	suite.db = &database.DB{DB: gormDB}

	// Migrate schema
	err = suite.db.AutoMigrate(&domain.User{})
	suite.Require().NoError(err)

	// Setup dependencies
	userRepo := repository.NewUserRepository(suite.db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// Setup Fiber app
	suite.app = fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Setup routes
	api := suite.app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "KBTG AI Backend Workshop is running!",
		})
	})

	users := api.Group("/users")
	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}

func (suite *APITestSuite) TearDownTest() {
	// Clean database between tests
	suite.db.Exec("DELETE FROM users")
}

func (suite *APITestSuite) TestHealthEndpoint() {
	// Act
	req := httptest.NewRequest("GET", "/api/v1/health", nil)
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(200, resp.StatusCode)

	// Parse response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)
	suite.Equal("ok", response["status"])
}

func (suite *APITestSuite) TestCreateUser() {
	// Arrange
	createReq := domain.CreateUserRequest{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		Points:         100,
	}

	body, err := json.Marshal(createReq)
	suite.Require().NoError(err)

	// Act
	req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(201, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)

	data := response["data"].(map[string]interface{})
	suite.Equal("John", data["first_name"])
	suite.Equal("Doe", data["last_name"])
	suite.Equal("john@example.com", data["email"])
	suite.NotEmpty(data["id"])
	suite.NotEmpty(data["membership_id"])
}

func (suite *APITestSuite) TestGetUsers() {
	// Arrange - Create test users
	users := []domain.User{
		{
			FirstName:      "John",
			LastName:       "Doe",
			Email:          "john@example.com",
			MembershipType: "Gold",
			MembershipID:   "LBK123456",
		},
		{
			FirstName:      "Jane",
			LastName:       "Smith",
			Email:          "jane@example.com",
			MembershipType: "Silver",
			MembershipID:   "LBK123457",
		},
	}

	for _, user := range users {
		err := suite.db.Create(&user).Error
		suite.Require().NoError(err)
	}

	// Act
	req := httptest.NewRequest("GET", "/api/v1/users", nil)
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(200, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)

	suite.Equal(float64(2), response["count"])

	data := response["data"].([]interface{})
	suite.Len(data, 2)
}

func (suite *APITestSuite) TestGetUserByID() {
	// Arrange - Create test user
	user := domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
	}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err)

	// Act
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/users/%d", user.ID), nil)
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(200, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)

	data := response["data"].(map[string]interface{})
	suite.Equal("John", data["first_name"])
	suite.Equal("john@example.com", data["email"])
}

func (suite *APITestSuite) TestGetUserByID_NotFound() {
	// Act
	req := httptest.NewRequest("GET", "/api/v1/users/999", nil)
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(404, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)
	suite.Equal("User not found", response["error"])
}

func (suite *APITestSuite) TestUpdateUser() {
	// Arrange - Create test user
	user := domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
		Points:         100,
	}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err)

	updateReq := domain.UpdateUserRequest{
		FirstName: "Jane",
		Points:    200,
	}

	body, err := json.Marshal(updateReq)
	suite.Require().NoError(err)

	// Act
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/v1/users/%d", user.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(200, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)

	data := response["data"].(map[string]interface{})
	suite.Equal("Jane", data["first_name"])
	suite.Equal("Doe", data["last_name"]) // unchanged
	suite.Equal(float64(200), data["points"])
}

func (suite *APITestSuite) TestDeleteUser() {
	// Arrange - Create test user
	user := domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
	}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err)

	// Act
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/users/%d", user.ID), nil)
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(200, resp.StatusCode)

	// Verify user is deleted
	var count int64
	suite.db.Model(&domain.User{}).Where("id = ?", user.ID).Count(&count)
	suite.Equal(int64(0), count)
}

func (suite *APITestSuite) TestCreateUser_ValidationError() {
	// Arrange - Invalid request (missing required fields)
	createReq := domain.CreateUserRequest{
		FirstName: "John",
		// LastName missing
		Email: "john@example.com",
	}

	body, err := json.Marshal(createReq)
	suite.Require().NoError(err)

	// Act
	req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(400, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)
	suite.Contains(response["error"], "required")
}

func (suite *APITestSuite) TestCreateUser_DuplicateEmail() {
	// Arrange - Create user first
	user := domain.User{
		FirstName:    "Existing",
		LastName:     "User",
		Email:        "duplicate@example.com",
		MembershipID: "LBK123456",
	}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err)

	// Try to create another user with same email
	createReq := domain.CreateUserRequest{
		FirstName: "New",
		LastName:  "User",
		Email:     "duplicate@example.com",
	}

	body, err := json.Marshal(createReq)
	suite.Require().NoError(err)

	// Act
	req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := suite.app.Test(req)

	// Assert
	suite.NoError(err)
	suite.Equal(400, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	suite.NoError(err)
	suite.Contains(response["error"], "already exists")
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
