package mocks

import (
	"github.com/stretchr/testify/mock"
	"kbtg.tech/ai-backend-workshop/internal/domain"
)

// MockUserRepository is a mock implementation of domain.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockUserUseCase is a mock implementation of domain.UserUseCase
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) GetAllUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserUseCase) GetUserByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUseCase) CreateUser(req domain.CreateUserRequest) (*domain.User, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUseCase) UpdateUser(id uint, req domain.UpdateUserRequest) (*domain.User, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUseCase) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
