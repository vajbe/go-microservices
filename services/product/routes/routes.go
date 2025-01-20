package routes

import (
	handlers "go-microservices/products/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	productHandler := handlers.NewProductHandler()

	// Define user service routes in a more concise way
	routes := []struct {
		method      string
		path        string
		handlerFunc func(http.ResponseWriter, *http.Request)
	}{
		/* 		{"GET", "/products", userHandler.GetUsers},
		   		{"GET", "/products/{id}", userHandler.GetUser}, */
		{"POST", "/products", productHandler.AddProduct},
		{"GET", "/products", productHandler.GetProducts},
	}

	// Register all routes in a loop to avoid repetition
	for _, route := range routes {
		router.HandleFunc(route.path, route.handlerFunc).Methods(route.method)
	}
}
