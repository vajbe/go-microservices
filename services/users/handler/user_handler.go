package handlers

import (
	"encoding/json"
	"fmt"
	"go-microservices/users/db"
	res "go-microservices/users/middleware"
	"go-microservices/users/types"

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
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "Users have been retrieved successfully.", users)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.AddUser(newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been added successfully.", resp)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.UpdateUser(newUser)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to update record: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been updated successfully.", newUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	var newUser types.User
	newUser.Id = idStr
	resultUser, err := db.GetUser(newUser)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to retrived record: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been retrived successfully.", resultUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	var newUser types.User
	newUser.Id = idStr
	err := db.DeleteUser(newUser)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to delete a record: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been deleted successfully.", newUser)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin types.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.Login(userLogin)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to login: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been logged in successfully.", resp)
}
