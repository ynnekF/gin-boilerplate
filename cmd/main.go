package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ynnekF/gin-boilerplate/middleware"
	"github.com/ynnekF/gin-boilerplate/internal"
	"go.uber.org/zap"
)

// Constant route path for health check endpoint
const HEALTH_PATH = "/health"

// getRouter initializes the Gin router with middleware and returns it.
func getRouter(logger *zap.SugaredLogger) *gin.Engine {
	router := gin.New()

	// Gin's built-in request logger and panic recovery middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Custom middleware
	router.Use(middleware.RequestLoggerMiddleware(logger))

	return router
}

// startApi registers the API routes and starts the HTTP server.
func startApi(router *gin.Engine, logger *zap.SugaredLogger) {
	// Register the health check endpoint
	router.GET(HEALTH_PATH, internal.GetHealth)

	// Start the server on port 8080 (blocking call)
	if err := router.Run(":8080"); err != nil {
    	logger.Fatalf("failed to start server: %v", err)
	}

}

func main() {
	// Initialize the zap logger
	logger := internal.CreateLogger()
	defer logger.Sync() // Flushes buffer, if any
	
	// Set Gin to release mode for production use
	gin.SetMode(gin.ReleaseMode)
	
	logger.Info("initializing router...")
	router := getRouter(logger)

	logger.Info("router initialized, starting server...")
	startApi(router, logger)
}