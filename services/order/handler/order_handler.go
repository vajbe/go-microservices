package handlers

import (
	"encoding/json"
	"go-microservices/order/db"
	"go-microservices/order/middleware"
	"go-microservices/order/types"

	"net/http"
)

type OrderHandler struct{}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

/*
	 func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
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
*/
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder types.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		middleware.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.CreateOrder(newOrder)
	if err != nil {
		middleware.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	middleware.Success(w, "Order has been placed successfully.", resp)
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
