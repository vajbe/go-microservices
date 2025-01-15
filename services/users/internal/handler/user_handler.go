package handlers

import (
	"encoding/json"
	"fmt"
	"go-microservices/users/internal/config"
	"go-microservices/users/internal/db"
	"go-microservices/users/internal/middleware/response"
	"go-microservices/users/internal/types"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n%+v", config.GetConfig())
	response.Success(w, "Users have been retrieved successfully.", nil)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.AddUser(newUser)
	response.Success(w, "User has been added successfully.", newUser)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	response.Error(w, fmt.Sprintf("User with ID: %s not found", idStr), http.StatusNotFound)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	response.Error(w, fmt.Sprintf("User with ID: %s not found", idStr), http.StatusNotFound)
}
