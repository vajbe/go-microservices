package handlers

import (
	"encoding/json"
	"fmt"
	"go-microservices/users/internal/middleware/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	response.Success(w, "Users have been retrieved successfully.", users)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, newUser)
	response.Success(w, "User has been added successfully.", newUser)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	for i, u := range users {
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, "Invalid ID Format.", http.StatusBadRequest)
			return
		}

		if u["id"] == id {
			var updatedData map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&updatedData)

			if err != nil {
				response.Error(w, "Error occurred while decoding updated data.", http.StatusInternalServerError)
				return
			}

			for key, value := range updatedData {
				u[key] = value
			}

			users[i] = u
			response.Success(w, fmt.Sprintf("User %d updated successfully.", id), updatedData)
			return
		}
	}
	response.Error(w, fmt.Sprintf("User with ID: %s not found", idStr), http.StatusNotFound)
}
