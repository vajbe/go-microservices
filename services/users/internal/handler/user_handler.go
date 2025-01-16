package handlers

import (
	"encoding/json"
	"fmt"
	"go-microservices/users/internal/db"
	"go-microservices/users/internal/middleware/response"
	"go-microservices/users/internal/types"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUsers()
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Success(w, "Users have been retrieved successfully.", users)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = db.AddUser(newUser)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Success(w, "User has been added successfully.", newUser)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.UpdateUser(newUser)
	if err != nil {
		response.Error(w, fmt.Sprintf("failed to update record: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	response.Success(w, "User has been updated successfully.", newUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	var newUser types.User
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUser.Id = id
	resultUser, err := db.GetUser(newUser)
	if err != nil {
		response.Error(w, fmt.Sprintf("failed to retrived record: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	response.Success(w, "User has been retrived successfully.", resultUser)
}
