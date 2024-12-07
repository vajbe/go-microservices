package routes

import (
	handlers "go-microservices/users/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userHandler := handlers.NewUserHandler()

	// Define user service routes
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
}
