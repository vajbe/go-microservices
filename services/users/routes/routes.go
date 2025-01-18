package routes

import (
	handlers "go-microservices/users/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userHandler := handlers.NewUserHandler()

	// Define user service routes in a more concise way
	routes := []struct {
		method      string
		path        string
		handlerFunc func(http.ResponseWriter, *http.Request)
	}{
		{"GET", "/users", userHandler.GetUsers},
		{"GET", "/users/{id}", userHandler.GetUser},
		{"POST", "/users", userHandler.AddUser},
		{"PUT", "/users/{id}", userHandler.UpdateUser},
		{"DELETE", "/users/{id}", userHandler.DeleteUser},
	}

	// Register all routes in a loop to avoid repetition
	for _, route := range routes {
		router.HandleFunc(route.path, route.handlerFunc).Methods(route.method)
	}
}
