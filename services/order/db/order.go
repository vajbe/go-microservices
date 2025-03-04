package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go-microservices/order/kafka"
	"go-microservices/order/types"
)

// CreateOrder inserts a new order with product details into the orders table
func CreateOrder(order types.Order) (types.Order, error) {
	pool := GetDBPool()

	// Serialize products into JSON format
	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return types.Order{}, fmt.Errorf("failed to marshal products: %w", err)
	}

	query := `
		INSERT INTO orders (user_id, order_status, total_amount, payment_status, products)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	var id string
	var createdAt int64

	// Insert the order with products JSON
	err = pool.QueryRow(context.Background(), query,
		order.UserID,
		order.OrderStatus,
		order.TotalAmount,
		order.PaymentStatus,
		productsJSON,
	).Scan(&id, &createdAt)
	if err != nil {
		return types.Order{}, fmt.Errorf("failed to insert order: %w", err)
	}

	// Return the created order
	createdOrder := types.Order{
		ID:            id,
		UserID:        order.UserID,
		OrderStatus:   order.OrderStatus,
		TotalAmount:   order.TotalAmount,
		PaymentStatus: order.PaymentStatus,
		Products:      order.Products,
		CreatedAt:     createdAt,
	}
	kManger := kafka.GetKafkaManger()
	kMsg, err := json.Marshal(createdOrder)
	if err != nil {
		fmt.Printf("Failed to marshal order message: %v", err)
	}
	err = kManger.Producer.Publish("order123", string(kMsg))
	if err != nil {
		fmt.Printf("Failed to publish message: %v", err)
	}
	return createdOrder, nil
}
