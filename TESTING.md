# Testing Guide

This project includes comprehensive unit tests for all layers of the Clean Architecture.

## Test Coverage

Current test coverage by layer:
- **Config**: 100.0% - Complete configuration testing
- **Handler**: 77.6% - HTTP handler testing with mocks
- **Repository**: 84.0% - Database integration testing
- **UseCase**: 72.3% - Business logic testing with mocks

## Test Structure

### 📁 Test Organization
```
internal/
├── config/
│   └── config_test.go          # Configuration tests
├── handler/
│   └── user_handler_test.go    # HTTP handler tests
├── repository/
│   └── user_repository_test.go # Database tests
├── usecase/
│   └── user_usecase_test.go    # Business logic tests
└── mocks/
    └── user_mocks.go           # Mock implementations
```

## Running Tests

### 🚀 Basic Testing
```bash
# Run all tests
go test ./internal/...

# Run tests with verbose output
go test -v ./internal/...

# Run tests for specific package
go test -v ./internal/usecase

# Using Makefile
make test
```

### 📊 Coverage Testing
```bash
# Generate coverage report
make test-coverage

# View coverage in browser
open coverage.html

# Check coverage percentage
go test -cover ./internal/...
```

## Test Types

### 1. Unit Tests (UseCase Layer)
```go
func TestUserUseCase_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := new(mocks.MockUserRepository)
    useCase := NewUserUseCase(mockRepo)
    
    req := domain.CreateUserRequest{...}
    mockRepo.On("GetByEmail", "john@example.com").Return(nil, errors.New("user not found"))
    mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)
    
    // Act
    result, err := useCase.CreateUser(req)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    mockRepo.AssertExpectations(t)
}
```

### 2. Integration Tests (Repository Layer)
```go
func (suite *UserRepositoryTestSuite) TestCreate() {
    // Uses in-memory SQLite database
    user := &domain.User{...}
    
    err := suite.repo.Create(user)
    
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), user.ID)
}
```

### 3. HTTP Handler Tests
```go
func TestUserHandler_GetUsers(t *testing.T) {
    // Arrange
    mockUseCase := new(mocks.MockUserUseCase)
    handler := NewUserHandler(mockUseCase)
    app := setupTestApp()
    
    expectedUsers := []domain.User{...}
    mockUseCase.On("GetAllUsers").Return(expectedUsers, nil)
    
    app.Get("/users", handler.GetUsers)
    
    // Act
    req := httptest.NewRequest("GET", "/users", nil)
    resp, err := app.Test(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, 200, resp.StatusCode)
}
```

## Testing Best Practices

### ✅ What We Test
1. **Business Logic** - All use case scenarios
2. **Error Handling** - Invalid inputs and edge cases
3. **Data Persistence** - Repository operations
4. **HTTP Responses** - Status codes and response format
5. **Validation** - Input validation and constraints

### 🎯 Testing Principles
1. **Arrange-Act-Assert** pattern
2. **Mock external dependencies** using interfaces
3. **Test both success and failure paths**
4. **Use descriptive test names**
5. **Test one thing at a time**

### 🛠 Tools Used
- **testify/assert** - Assertions and test utilities
- **testify/mock** - Mock generation and validation
- **testify/suite** - Test suite for setup/teardown
- **httptest** - HTTP testing utilities
- **In-memory SQLite** - Database testing without external deps

## Mock Usage

### Creating Mocks
```go
// Repository mock
mockRepo := new(mocks.MockUserRepository)
mockRepo.On("GetByID", uint(1)).Return(expectedUser, nil)

// UseCase mock  
mockUseCase := new(mocks.MockUserUseCase)
mockUseCase.On("CreateUser", req).Return(expectedUser, nil)
```

### Mock Validation
```go
// Verify all expected calls were made
mockRepo.AssertExpectations(t)

// Verify specific calls
mockRepo.AssertCalled(t, "GetByID", uint(1))
```

## Test Data

### Test User Creation
```go
func createTestUser() *domain.User {
    return &domain.User{
        FirstName:      "John",
        LastName:       "Doe", 
        Email:          "john@example.com",
        Phone:          "123-456-7890",
        MembershipType: "Gold",
        MembershipID:   "LBK123456",
        Points:         100,
    }
}
```

## Continuous Integration

### Pre-commit Testing
```bash
# Run full test suite before commit
make test

# Check test coverage
make test-coverage

# Ensure code quality
make fmt vet
```

## Troubleshooting

### Common Issues
1. **Mock not matching** - Check method signatures and parameters
2. **Database locks** - Ensure proper test cleanup
3. **Race conditions** - Use -race flag: `go test -race ./...`

### Debugging Tests
```bash
# Run specific test with verbose output
go test -v -run TestUserUseCase_CreateUser ./internal/usecase

# Run with race detection
go test -race ./internal/...

# Run with coverage and race detection
go test -race -cover ./internal/...
```

## Adding New Tests

### 1. Create Test File
```bash
# For new package
touch internal/newpackage/newpackage_test.go
```

### 2. Add Test Function
```go
func TestNewFunction(t *testing.T) {
    // Arrange
    
    // Act
    
    // Assert
}
```

### 3. Update Mocks (if needed)
```go
// Add new method to mock in internal/mocks/
func (m *MockInterface) NewMethod(param string) error {
    args := m.Called(param)
    return args.Error(0)
}
```

## Performance Testing

### Benchmarks
```bash
# Run benchmarks
go test -bench=. ./internal/...

# Run specific benchmark
go test -bench=BenchmarkCreateUser ./internal/usecase
```

### Example Benchmark
```go
func BenchmarkUserUseCase_CreateUser(b *testing.B) {
    mockRepo := new(mocks.MockUserRepository)
    useCase := NewUserUseCase(mockRepo)
    
    req := domain.CreateUserRequest{...}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        useCase.CreateUser(req)
    }
}
```
