package main

import (
	"log"
	"net/http"

	"go-microservices/users/config"
	"go-microservices/users/db"
	"go-microservices/users/middleware/logging"
	"go-microservices/users/routes"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var (
	REDIS_CLIENT   *redis.Client
	SERVICE_CONFIG config.Config
)

func main() {
	// Load configuration
	SERVICE_CONFIG = config.Load()

	// Initialize Table and dbs
	db.InitializeDb(SERVICE_CONFIG)

	//	Initialize Redis
	db.InitRedis(SERVICE_CONFIG)

	// Initialize router
	router := mux.NewRouter()

	// Apply middlewares
	router.Use(logging.LoggingMiddleware)

	// Register routes
	routes.RegisterUserRoutes(router)

	// Start server
	log.Printf("User Service running on port %s", SERVICE_CONFIG.Port)
	log.Fatal(http.ListenAndServe(":"+SERVICE_CONFIG.Port, router))
}
