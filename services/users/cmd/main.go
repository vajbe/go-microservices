package main

import (
	"log"
	"net/http"

	"go-microservices/users/internal/config"
	"go-microservices/users/internal/middleware/logging"
	"go-microservices/users/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize router
	router := mux.NewRouter()

	// Apply middlewares
	router.Use(logging.LoggingMiddleware)

	// Register routes
	routes.RegisterUserRoutes(router)

	// Start server
	log.Printf("User Service running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
