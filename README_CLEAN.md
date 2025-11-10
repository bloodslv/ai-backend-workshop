# KBTG AI Backend Workshop - Clean Architecture

A REST API backend built with Go, Fiber framework, and Clean Architecture principles.

## Architecture Overview

This project follows Clean Architecture principles with clear separation of concerns:

```
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── domain/           # Domain entities and interfaces
│   ├── usecase/          # Business logic
│   ├── repository/       # Data access layer
│   ├── handler/          # HTTP handlers (controllers)
│   └── config/           # Configuration
├── pkg/                   # Public packages
│   └── database/         # Database utilities
├── public/               # Static files
└── main.go               # Application entry point
```

### Layers

1. **Domain Layer** (`internal/domain/`)
   - Contains business entities (User)
   - Defines repository and use case interfaces
   - No dependencies on external packages

2. **Use Case Layer** (`internal/usecase/`)
   - Contains business logic
   - Implements domain interfaces
   - Coordinates between different repositories

3. **Repository Layer** (`internal/repository/`)
   - Implements data access
   - Handles database operations
   - Implements domain repository interfaces

4. **Handler Layer** (`internal/handler/`)
   - HTTP request/response handling
   - Input validation
   - Calls use cases

5. **Infrastructure** (`pkg/`)
   - Database connection and utilities
   - External service integrations

## Prerequisites

- Go 1.17 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd ai-backend-workshop
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:3000`

## Configuration

Environment variables:

- `PORT` - Server port (default: 3000)
- `DB_PATH` - SQLite database file path (default: users.db)
- `APP_NAME` - Application name (default: KBTG AI Backend Workshop)
- `DEBUG` - Debug mode (default: false)

## API Endpoints

### Health Check
- **GET** `/api/v1/health` - Check if the server is running

### Hello World
- **GET** `/api/v1/hello?name=YourName` - Simple greeting endpoint

### Users CRUD
- **GET** `/api/v1/users` - Get all users
- **GET** `/api/v1/users/:id` - Get user by ID
- **POST** `/api/v1/users` - Create a new user
- **PUT** `/api/v1/users/:id` - Update a user
- **DELETE** `/api/v1/users/:id` - Delete a user

### User Model

```json
{
  "id": 1,
  "first_name": "สมชาย",
  "last_name": "ใจดี",
  "email": "somchai@example.com",
  "phone": "081-234-5678",
  "membership_type": "Gold",
  "membership_id": "LBK001234",
  "join_date": "2023-11-10T13:39:37.123Z",
  "points": 15420,
  "created_at": "2024-11-10T13:39:37.123Z",
  "updated_at": "2024-11-10T13:39:37.123Z"
}
```

## Example Usage

### Create a user
```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "phone": "081-234-5678",
    "membership_type": "Gold",
    "points": 1000
  }'
```

### Get all users
```bash
curl http://localhost:3000/api/v1/users
```

### Update a user
```bash
curl -X PUT http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "points": 2000
  }'
```

## Project Structure Benefits

### 🏗️ Clean Architecture Principles

1. **Independence of Frameworks** - The architecture doesn't depend on external frameworks
2. **Testable** - Business rules can be tested without external dependencies
3. **Independent of UI** - Easy to change UI without changing business rules
4. **Independent of Database** - Can switch databases without affecting business logic
5. **Independent of External Services** - Business rules don't know about external interfaces

### 🔧 Maintainability

- **Single Responsibility** - Each layer has one reason to change
- **Dependency Inversion** - High-level modules don't depend on low-level modules
- **Interface Segregation** - Clients don't depend on interfaces they don't use

### 🚀 Scalability

- Easy to add new features
- Simple to modify existing functionality
- Clear separation makes team collaboration easier

## Development

### Adding New Features

1. **Define Domain Entity** in `internal/domain/`
2. **Create Repository Interface** in `internal/domain/`
3. **Implement Repository** in `internal/repository/`
4. **Create Use Case Interface** in `internal/domain/`
5. **Implement Use Case** in `internal/usecase/`
6. **Create HTTP Handler** in `internal/handler/`
7. **Add Routes** in `main.go`

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o app main.go
```

## Built With

- [Go](https://golang.org/) - Programming language
- [Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [SQLite](https://sqlite.org/) - Database

## Contributing

1. Follow Clean Architecture principles
2. Keep layers independent
3. Write tests for business logic
4. Document public APIs
5. Follow Go conventions
