package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go-microservices/products/db"
	res "go-microservices/products/middleware"
	"go-microservices/products/types"

	"net/http"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	sortBy := r.URL.Query().Get("sort_by")
	orderBy := r.URL.Query().Get("order_by")
	name := r.URL.Query().Get("name")
	if orderBy == "" {
		orderBy = "desc"
	}

	if sortBy == "" {
		sortBy = "created_at"
	}

	if limitStr == "" || offsetStr == "" {
		res.Error(w, fmt.Errorf("error while decoding query params").Error(), http.StatusBadRequest)
		return
	}

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	products, err := db.GetProducts(limit, offset, sortBy, orderBy, name)
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "Products have been retrieved successfully.", products)
}

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
