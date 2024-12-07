package main

import (
	"log"
	"net/http"

	"go-microservices/users/internal/config"
	"go-microservices/users/internal/middleware"
	"go-microservices/users/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize router
	router := mux.NewRouter()

	// Apply middlewares
	router.Use(middleware.LoggingMiddleware)

	// Register routes
	routes.RegisterUserRoutes(router)

	// Start server
	log.Printf("User Service running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
