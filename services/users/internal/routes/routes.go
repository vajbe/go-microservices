package routes

import (
	handlers "go-microservices/users/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userHandler := handlers.NewUserHandler()

	// Define user service routes
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users", userHandler.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
}
