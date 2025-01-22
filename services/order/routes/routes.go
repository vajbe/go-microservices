package routes

import (
	handlers "go-microservices/order/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	orderHandler := handlers.NewOrderHandler()

	// Define user service routes in a more concise way
	routes := []struct {
		method      string
		path        string
		handlerFunc func(http.ResponseWriter, *http.Request)
	}{
		{"POST", "/order", orderHandler.CreateOrder},
	}

	// Register all routes in a loop to avoid repetition
	for _, route := range routes {
		router.HandleFunc(route.path, route.handlerFunc).Methods(route.method)
	}
}
