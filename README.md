# AI Backend Workshop - Go + Fiber Clean Architecture

A RESTful API backend built with Go and Fiber framework following Clean Architecture principles.

## 🚀 Features

- **Clean Architecture**: Separation of concerns with clear layer boundaries
- **RESTful API**: Complete CRUD operations for user management
- **Database**: SQLite with GORM ORM and auto-migration
- **Testing**: Comprehensive unit tests and integration tests (73.7% coverage)
- **Configuration**: Environment-based configuration management
- **Mock Testing**: Isolated unit testing with mocks
- **Validation**: Request validation and error handling

## 📁 Project Structure

```
├── internal/
│   ├── config/          # Configuration management
│   ├── domain/          # Domain entities and interfaces
│   ├── handler/         # HTTP handlers (Presentation layer)
│   ├── mocks/          # Mock implementations for testing
│   ├── repository/      # Data access layer
│   └── usecase/        # Business logic layer
├── pkg/
│   └── database/       # Database connection and utilities
├── tests/              # Integration tests
├── main.go            # Application entry point
├── go.mod             # Go module dependencies
├── Makefile          # Build and development commands
└── README.md         # Project documentation
```

## 🛠️ Tech Stack

- **Language**: Go 1.17+
- **Web Framework**: Fiber v2.52.9
- **Database**: SQLite
- **ORM**: GORM v1.25.12
- **Testing**: testify/assert, testify/mock, testify/suite
- **Configuration**: Environment variables

## 🔧 Setup and Installation

### Prerequisites

- Go 1.17 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd ai-backend-workshop
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set environment variables (optional):
```bash
export PORT=3000
export DB_PATH=./test.db
export ENV=development
```

4. Run the application:
```bash
make run
# or
go run main.go
```

## 🚀 Available Commands

```bash
# Run the application
make run

# Run all tests
make test

# Run tests with coverage
make test-coverage

# View coverage report in browser
make coverage-html

# Clean build artifacts
make clean

# Build the application
make build
```

## 📊 API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### User Management
- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create new user
- `PUT /api/users/:id` - Update user by ID
- `DELETE /api/users/:id` - Delete user by ID

## Example Usage

### Create User
```bash
curl -X POST http://localhost:3000/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
  }'
```

### Get All Users
```bash
curl http://localhost:3000/api/users
```

### Update User
```bash
curl -X PUT http://localhost:3000/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "email": "johnsmith@example.com",
    "age": 31
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:3000/api/users/1
```

## 🧪 Testing

The project includes comprehensive testing with different types of tests:

### Unit Tests
- **Handler Tests**: HTTP handler testing with mocked dependencies
- **UseCase Tests**: Business logic testing with mocked repositories
- **Repository Tests**: Database operations testing with in-memory SQLite
- **Config Tests**: Configuration management testing

### Integration Tests
- **API Tests**: Full API endpoint testing with real database operations
- **Database Integration**: Complete CRUD operations testing

### Run Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Generate coverage report
make coverage-html

# Run specific test package
go test -v ./internal/handler/...
go test -v ./internal/usecase/...
go test -v ./internal/repository/...
go test -v ./tests/...
```

### Test Coverage

Current test coverage: **73.7%**

- Config: 100.0%
- Handler: 77.6%
- Repository: 84.0%
- UseCase: 72.3%

## 🏗️ Architecture

This project follows Clean Architecture principles:

### Layers

1. **Domain Layer** (`internal/domain/`): Contains business entities, interfaces, and DTOs
2. **Use Case Layer** (`internal/usecase/`): Contains business logic and application rules
3. **Repository Layer** (`internal/repository/`): Contains data access logic
4. **Handler Layer** (`internal/handler/`): Contains HTTP handlers and routing

### Dependency Direction

```
Handler → UseCase → Repository → Database
  ↓         ↓         ↓
Domain ← Domain ← Domain
```

### Key Principles

- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Single Responsibility**: Each layer has a single, well-defined responsibility
- **Interface Segregation**: Dependencies are injected through interfaces
- **Testability**: Each layer can be tested in isolation using mocks

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `3000` | Server port |
| `DB_PATH` | `./test.db` | SQLite database file path |
| `ENV` | `development` | Environment mode |

## 📝 Development

### Adding New Features

1. **Add Domain Entity**: Update `internal/domain/` with new entities and interfaces
2. **Implement Use Case**: Add business logic in `internal/usecase/`
3. **Implement Repository**: Add data access logic in `internal/repository/`
4. **Add Handler**: Create HTTP endpoints in `internal/handler/`
5. **Write Tests**: Add unit tests and integration tests
6. **Update Documentation**: Update README and API documentation

### Database Migrations

The application automatically creates and migrates the database schema on startup. To add new models:

1. Define the model in `internal/domain/`
2. Add migration in `pkg/database/database.go`
3. Update seed data if needed

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 🙏 Acknowledgments

- [Fiber](https://gofiber.io/) - Fast HTTP framework for Go
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [Testify](https://github.com/stretchr/testify) - Testing toolkit for Go
- Clean Architecture concepts by Robert C. Martin

### Create a user
```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Update a user
```bash
curl -X PUT http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane@example.com"}'
```

### Delete a user
```bash
curl -X DELETE http://localhost:3000/api/v1/users/1
```

## Environment Variables

- `PORT` - Server port (default: 3000)

## Project Structure

```
.
├── main.go           # Main application file
├── public/          # Static files
│   └── index.html   # Welcome page
├── go.mod           # Go module file
├── go.sum           # Go module checksums
└── README.md        # This file
```

## Features

- RESTful API with CRUD operations
- JSON responses
- CORS enabled
- Request logging
- Error recovery middleware
- Static file serving
- Environment configuration

## Built With

- [Go](https://golang.org/) - Programming language
- [Fiber](https://gofiber.io/) - Web framework
- [Fasthttp](https://github.com/valyala/fasthttp) - HTTP engine

## Development

To modify the application:

1. Edit `main.go` to add new routes or modify existing ones
2. Add static files to the `public/` directory
3. Use `go run main.go` to restart the server

## Production Deployment

1. Build the application:
```bash
go build -o app main.go
```

2. Run the built binary:
```bash
./app
```

Or set the PORT environment variable:
```bash
PORT=8080 ./app
```
