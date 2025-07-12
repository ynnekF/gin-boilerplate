# GO Gin Boilerplate
---
A minimal, production-ready REST API boilerplate built with [Go](https://golang.org) and [Gin](https://github.com/gin-gonic/gin). This project sets up a clean foundation for building scalable backend services using structured logging, modular routing, and middleware.

## Features
- Lighweight HTTP router using GIN
- Structured logging with [Zap](https://github.com/uber-go/zap)
- Health check endpoint
- Request/Response Logging Middleware

## Structure
### `/api`
- OpenAPI and swagger documentation

### `/bin`
- Build objects

### `/cmd`
- Main application and entrypoint for project

### `/internal`
- Internal logic

### `/middleware`
- App middleware

### `/models`
- Standard types/structs for processing

## How to Run
```bash
# Clone the repository
git clone https://github.com/yourusername/gin-boilerplate.git
cd gin-boilerplate

# Initialize Go modules (if needed)
go mod tidy

# Run the service
go run main.go
```

### `make run`
- Executes `main.go`
### `make build`
- Creates executable
### `make clean`
- Removes build objects

## Requirements

- Go 1.20+
- Git

## Other Commands
```bash
go fmt ./...           # Format code
go vet ./...           # Static analysis
go test ./...          # Run tests
```