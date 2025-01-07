package handlers

import (
	"encoding/json"
	"go-microservices/users/internal/middleware/response"
	"net/http"
)

type UserHandler struct{}

var (
	users []map[string]interface{}
)

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func init() {
	users = []map[string]interface{}{
		{"id": 1, "name": "John Doe"},
		{"id": 2, "name": "Jane Doe"},
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, newUser)

	response.Success(w, "User has been added successfully", newUser)
}
