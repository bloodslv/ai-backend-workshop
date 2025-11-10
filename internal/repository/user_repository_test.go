package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kbtg.tech/ai-backend-workshop/internal/domain"
	"kbtg.tech/ai-backend-workshop/pkg/database"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db   *database.DB
	repo domain.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	// Create in-memory SQLite database for testing
	gormDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Require().NoError(err)

	suite.db = &database.DB{DB: gormDB}

	// Migrate the schema
	err = suite.db.AutoMigrate(&domain.User{})
	suite.Require().NoError(err)

	suite.repo = NewUserRepository(suite.db)
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	// Arrange
	user := &domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
		Points:         100,
	}

	// Act
	err := suite.repo.Create(user)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), user.ID)
	assert.NotZero(suite.T(), user.CreatedAt)
	assert.NotZero(suite.T(), user.UpdatedAt)
}

func (suite *UserRepositoryTestSuite) TestGetByID() {
	// Arrange
	user := &domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
		Points:         100,
	}
	err := suite.repo.Create(user)
	suite.Require().NoError(err)

	// Act
	result, err := suite.repo.GetByID(user.ID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.FirstName, result.FirstName)
	assert.Equal(suite.T(), user.LastName, result.LastName)
	assert.Equal(suite.T(), user.Email, result.Email)
}

func (suite *UserRepositoryTestSuite) TestGetByID_NotFound() {
	// Act
	result, err := suite.repo.GetByID(999)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), "user not found", err.Error())
}

func (suite *UserRepositoryTestSuite) TestGetByEmail() {
	// Arrange
	user := &domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
		Points:         100,
	}
	err := suite.repo.Create(user)
	suite.Require().NoError(err)

	// Act
	result, err := suite.repo.GetByEmail("john@example.com")

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.FirstName, result.FirstName)
	assert.Equal(suite.T(), user.Email, result.Email)
}

func (suite *UserRepositoryTestSuite) TestGetByEmail_NotFound() {
	// Act
	result, err := suite.repo.GetByEmail("notfound@example.com")

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), "user not found", err.Error())
}

func (suite *UserRepositoryTestSuite) TestGetAll() {
	// Arrange
	users := []*domain.User{
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
		err := suite.repo.Create(user)
		suite.Require().NoError(err)
	}

	// Act
	result, err := suite.repo.GetAll()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 2)
}

func (suite *UserRepositoryTestSuite) TestUpdate() {
	// Arrange
	user := &domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		Phone:          "123-456-7890",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
		Points:         100,
	}
	err := suite.repo.Create(user)
	suite.Require().NoError(err)

	// Act
	user.FirstName = "Jane"
	user.Points = 200
	err = suite.repo.Update(user)

	// Assert
	assert.NoError(suite.T(), err)

	// Verify update
	updated, err := suite.repo.GetByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Jane", updated.FirstName)
	assert.Equal(suite.T(), 200, updated.Points)
}

func (suite *UserRepositoryTestSuite) TestDelete() {
	// Arrange
	user := &domain.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@example.com",
		MembershipType: "Gold",
		MembershipID:   "LBK123456",
	}
	err := suite.repo.Create(user)
	suite.Require().NoError(err)

	// Act
	err = suite.repo.Delete(user.ID)

	// Assert
	assert.NoError(suite.T(), err)

	// Verify deletion
	_, err = suite.repo.GetByID(user.ID)
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestDelete_NotFound() {
	// Act
	err := suite.repo.Delete(999)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "user not found", err.Error())
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
