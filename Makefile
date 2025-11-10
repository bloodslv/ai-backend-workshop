.PHONY: run build clean test fmt vet deps lint

# Default target
all: deps fmt vet test build

# Run the application in development mode
run:
	go run main.go

# Build the application
build:
	go build -o app main.go

# Clean build artifacts
clean:
	rm -f app users.db
	go clean

# Run tests
test:
	go test ./internal/...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./internal/...
	go tool cover -html=coverage.out -o coverage.html

# Format Go code
fmt:
	go fmt ./...

# Vet Go code
vet:
	go vet ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Download dependencies
deps:
	go mod download
	go mod tidy

# Install development dependencies
dev-deps:
	go install github.com/air-verse/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run with hot reload (requires air to be installed)
dev:
	air

# Build for production (Linux)
build-prod:
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app main.go

# Build for production (current OS)
build-local:
	go build -ldflags="-s -w" -o app main.go

# Run with specific port
run-port:
	PORT=8080 go run main.go

# Generate mocks (requires mockgen)
generate:
	go generate ./...

# Database operations
db-reset:
	rm -f users.db

# Run security scan
security:
	gosec ./...

# Check for updates
updates:
	go list -u -m all

# Help
help:
	@echo "Available commands:"
	@echo "  run           - Run the application in development mode"
	@echo "  build         - Build the application"
	@echo "  build-prod    - Build for production (Linux)"
	@echo "  build-local   - Build for production (current OS)"
	@echo "  clean         - Clean build artifacts and database"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  fmt           - Format Go code"
	@echo "  vet           - Vet Go code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  deps          - Download and tidy dependencies"
	@echo "  dev-deps      - Install development dependencies"
	@echo "  dev           - Run with hot reload (requires air)"
	@echo "  run-port      - Run with PORT=8080"
	@echo "  generate      - Generate mocks and other generated code"
	@echo "  db-reset      - Reset database"
	@echo "  security      - Run security scan"
	@echo "  updates       - Check for dependency updates"
	@echo "  help          - Show this help message"
