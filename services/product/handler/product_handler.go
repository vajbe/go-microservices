package handlers

import (
	"encoding/json"

	"go-microservices/products/db"
	res "go-microservices/products/middleware"
	"go-microservices/products/types"

	"net/http"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

/* func (h *ProductHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUsers()
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "Products have been retrieved successfully.", users)
} */

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var newrProduct types.Product
	err := json.NewDecoder(r.Body).Decode(&newrProduct)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.AddProduct(newrProduct)
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "Product has been added successfully.", resp)
}

/* func (h *ProductHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	var newProduct types.Product
	newProduct.Id = idStr
	resultUser, err := db.GetUser(newProduct)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to retrived record: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "Product has been retrived successfully.", resultUser)
}
*/
