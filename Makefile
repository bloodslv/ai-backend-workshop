.PHONY: run build clean test fmt vet deps

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
	rm -f app
	go clean

# Run tests
test:
	go test ./...

# Format Go code
fmt:
	go fmt ./...

# Vet Go code
vet:
	go vet ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Install development dependencies
dev-deps:
	go install github.com/air-verse/air@latest

# Run with hot reload (requires air to be installed)
dev:
	air

# Build for production
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

# Run with specific port
run-port:
	PORT=8080 go run main.go

# Help
help:
	@echo "Available commands:"
	@echo "  run        - Run the application in development mode"
	@echo "  build      - Build the application"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  fmt        - Format Go code"
	@echo "  vet        - Vet Go code"
	@echo "  deps       - Download and tidy dependencies"
	@echo "  dev-deps   - Install development dependencies (air for hot reload)"
	@echo "  dev        - Run with hot reload (requires air)"
	@echo "  build-prod - Build for production (Linux binary)"
	@echo "  run-port   - Run with PORT=8080"
	@echo "  help       - Show this help message"
