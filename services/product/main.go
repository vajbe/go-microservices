package main

import (
	"go-microservices/products/config"
	"go-microservices/products/db"
	"go-microservices/products/middleware"
	"go-microservices/products/routes"
	"log"
	"net/http"

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
	router.Use(middleware.LoggingMiddleware)

	// Register routes
	routes.RegisterProductRoutes(router)

	// Start server
	log.Printf("Product Service running on port %s", SERVICE_CONFIG.Port)
	log.Fatal(http.ListenAndServe(":"+SERVICE_CONFIG.Port, router))
}
