package db

import (
	"context"
	"fmt"
	"go-microservices/products/types"
)

func AddProduct(product types.Product) (types.ProductResponse, error) {
	pool := GetDBPool()
	query := `INSERT INTO products (name, description, price, stock, category) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	var id string
	var createdAt int64

	err := pool.QueryRow(context.Background(), query, product.Name, product.Description, product.Price, product.Stock, product.Category).Scan(&id, &createdAt)
	if err != nil {
		return types.ProductResponse{Id: product.Id}, fmt.Errorf("failed to insert record: %w", err)
	}
	response := types.ProductResponse{
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Id:          id,
		Category:    product.Category,
		Price:       product.Price,
	}
	return response, nil
}

func GetProducts() ([]types.ProductResponse, error) {
	pool := GetDBPool()
	rows, err := pool.Query(context.Background(), "SELECT ID, NAME, CATEGORY, PRICE, STOCK, DESCRIPTION FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	var products []types.ProductResponse
	for rows.Next() {
		var product types.ProductResponse
		// Scan each row into the User struct
		err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Price, &product.Stock, &product.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		products = append(products, product)
	}
	// Check for errors after the loop
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return products, nil
}

/*func GetUser(user types.User) (types.User, error) {
	pool := GetDBPool()
	query := "SELECT * FROM users where id=$1"
	var resultUser types.User
	err := pool.QueryRow(context.Background(), query, user.Id).Scan(&resultUser.Id, &resultUser.Name, &resultUser.Email, &resultUser.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return resultUser, fmt.Errorf("no record found for id: %s", user.Id)
		}
		return resultUser, fmt.Errorf("failed to execute query: %w", err)
	}

	return resultUser, nil
}
*/
