# GitHub Copilot Instructions

## Project Overview

This is a **Go + Fiber Clean Architecture** backend project for AI Backend Workshop. The system implements RESTful APIs with clean separation of concerns following Domain-Driven Design principles.

## Code Style Guidelines

### Go Code Style
- Follow **Go standard formatting** using `gofmt` and `goimports`
- Use **camelCase** for variables and functions, **PascalCase** for exported types
- Keep **line length under 120 characters** for better readability
- Use **meaningful variable names** - avoid single letters except for loop counters
- Add **comprehensive comments** for exported functions and types
- Use **struct tags** consistently: `json:"field_name" gorm:"constraint" validate:"rule"`

### Naming Conventions
```go
// ✅ Good examples
type UserRepository interface { }
type CreateUserRequest struct { }
func (uc *userUseCase) GetAllUsers() ([]User, error) { }

// ❌ Bad examples  
type userrepo interface { }
type createreq struct { }
func (u *usecase) get() ([]user, error) { }
```

### Error Handling
- Always return **meaningful error messages** with context
- Use **fmt.Errorf** with verb `%w` for error wrapping
- Log errors at appropriate levels: `log.Error`, `log.Warn`, `log.Info`
- Return **appropriate HTTP status codes** in handlers

```go
// ✅ Good error handling
if err != nil {
    return nil, fmt.Errorf("failed to get user by ID %d: %w", id, err)
}

// ❌ Poor error handling  
if err != nil {
    return nil, err
}
```

## Architecture Guidelines

### Clean Architecture Layers

This project follows **Clean Architecture** with these layers:

```
┌─────────────────┐
│   Handler       │  ← HTTP routing, request validation, response formatting
│  (Presentation) │
└─────────────────┘
         │
         v
┌─────────────────┐
│   Use Case      │  ← Business logic, orchestration, validation
│ (Application)   │
└─────────────────┘
         │
         v  
┌─────────────────┐
│   Repository    │  ← Data access, database operations
│(Infrastructure) │
└─────────────────┘
         │
         v
┌─────────────────┐
│    Domain       │  ← Entities, interfaces, business rules
│   (Enterprise)  │
└─────────────────┘
```

### Layer Responsibilities

#### Domain Layer (`internal/domain/`)
- Define **entities**, **interfaces**, and **business rules**
- Contains **no dependencies** on external frameworks
- Define **repository interfaces** and **use case interfaces**
- Include **request/response DTOs**

#### Use Case Layer (`internal/usecase/`)
- Implement **business logic** and **application rules**
- Orchestrate calls between repositories
- Handle **validation** and **error scenarios**
- **Depend only on domain interfaces**

#### Repository Layer (`internal/repository/`)
- Implement **data access logic** using GORM
- Handle **database transactions** and **queries**
- Implement **domain repository interfaces**
- Convert between **database models** and **domain entities**

#### Handler Layer (`internal/handler/`)
- Handle **HTTP requests** and **responses**
- Perform **request validation** and **parameter parsing**
- Convert between **HTTP formats** and **domain objects**
- Return **appropriate status codes** and **error responses**

### Dependency Injection Pattern
- Use **constructor functions** with `New` prefix
- Inject **interfaces, not concrete types**
- Follow **dependency inversion principle**

```go
// ✅ Correct dependency injection
func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
    return &userUseCase{userRepo: userRepo}
}

// ❌ Wrong - depends on concrete type
func NewUserUseCase(userRepo *repository.UserRepository) domain.UserUseCase {
    return &userUseCase{userRepo: userRepo}
}
```

## What NOT to Do

### ❌ Architecture Violations
- **DO NOT** access database directly from handlers
- **DO NOT** put business logic in repository layer  
- **DO NOT** import upper layers from lower layers
- **DO NOT** use concrete types in interfaces
- **DO NOT** mix HTTP concerns with business logic

### ❌ Code Anti-patterns
- **DO NOT** use global variables for shared state
- **DO NOT** ignore errors or return nil without checking
- **DO NOT** use `panic()` for normal error handling
- **DO NOT** write functions longer than 50 lines
- **DO NOT** use magic numbers or hardcoded strings

### ❌ Database Anti-patterns
- **DO NOT** write raw SQL queries (use GORM methods)
- **DO NOT** forget to handle database transactions
- **DO NOT** expose GORM models outside repository layer
- **DO NOT** perform N+1 queries (use preloading)

### ❌ Testing Anti-patterns
- **DO NOT** test implementation details
- **DO NOT** write tests without assertions
- **DO NOT** use real database in unit tests
- **DO NOT** forget to test error scenarios

## File Organization

```
internal/
├── config/          # Configuration management
│   ├── config.go    # Environment configuration
│   └── config_test.go
├── domain/          # Business entities & interfaces  
│   ├── user.go      # User domain model & interfaces
│   └── transfer.go  # Transfer domain model & interfaces
├── usecase/         # Business logic implementation
│   ├── user_usecase.go
│   ├── user_usecase_test.go
│   └── transfer_usecase.go
├── repository/      # Data access layer
│   ├── user_repository.go
│   ├── user_repository_test.go  
│   └── transfer_repository.go
├── handler/         # HTTP handlers
│   ├── user_handler.go
│   ├── user_handler_test.go
│   └── transfer_handler.go
└── mocks/           # Mock implementations for testing
    └── user_mocks.go
```

## Testing Guidelines

### Unit Testing
- Write **table-driven tests** for multiple scenarios
- Use **testify/assert** for assertions  
- Use **testify/mock** for mocking dependencies
- Test **both success and error cases**
- Aim for **>80% code coverage**

```go
// ✅ Good test structure
func TestUserUseCase_CreateUser(t *testing.T) {
    tests := []struct {
        name    string
        req     domain.CreateUserRequest  
        want    *domain.User
        wantErr bool
    }{
        {
            name: "successful creation",
            req:  domain.CreateUserRequest{...},
            want: &domain.User{...}, 
            wantErr: false,
        },
        {
            name: "validation error",
            req:  domain.CreateUserRequest{},
            want: nil,
            wantErr: true, 
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Integration Testing
- Test **full API endpoints** with real HTTP calls
- Use **in-memory databases** for testing
- Test **complete workflows** end-to-end
- Verify **HTTP status codes** and **response formats**

## API Design Guidelines

### RESTful Conventions
- Use **proper HTTP methods**: GET, POST, PUT, DELETE
- Use **consistent URL patterns**: `/api/users/:id`
- Return **appropriate status codes**: 200, 201, 400, 404, 500
- Use **JSON for request/response bodies**

### Response Format
```go
// ✅ Success response
{
    "id": 1,
    "name": "John Doe", 
    "email": "john@example.com"
}

// ✅ Error response
{
    "error": "VALIDATION_ERROR",
    "message": "email is required"
}
```

### Request Validation
- Validate **required fields**
- Check **data formats** (email, phone, etc.)
- Enforce **business rules** (positive amounts, valid statuses)
- Return **clear validation error messages**

## Performance Guidelines

### Database Optimization
- Use **proper indexes** for frequently queried fields
- Implement **pagination** for large datasets  
- Use **GORM preloading** to avoid N+1 queries
- Consider **database connection pooling**

### Memory Management
- **Close database connections** properly
- **Limit response payload size** with pagination
- Use **streaming** for large file operations
- Implement **proper garbage collection**

## Security Guidelines

### Input Validation
- **Sanitize all user inputs**
- **Validate data types** and **ranges**
- **Prevent SQL injection** (GORM handles this)
- **Limit request payload size**

### Error Handling
- **Don't expose internal errors** to clients
- **Log security events** appropriately
- **Use HTTPS** in production
- **Implement rate limiting**

---

*These instructions help GitHub Copilot generate code that follows our project's Clean Architecture patterns, Go best practices, and maintains consistency across the codebase.*