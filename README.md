# KBTG AI Backend Workshop

A simple REST API backend built with Go and Fiber framework.

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

## Example Usage

### Get all users
```bash
curl http://localhost:3000/api/v1/users
```

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
