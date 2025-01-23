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

func GetProducts(limit int, offset int, sortBy string, orderBy string, name string) (types.ProductListingResponse, error) {
	pool := GetDBPool()
	query := fmt.Sprintf("SELECT ID, NAME, CATEGORY, PRICE, STOCK, DESCRIPTION FROM products "+
		"WHERE NAME LIKE $3"+
		" ORDER BY %s %s LIMIT $1 OFFSET $2", sortBy, orderBy)

	if name == "" {
		name = "%"
	} else {
		name = "%" + name + "%"
	}

	rows, err := pool.Query(context.Background(), query, limit, offset, name)
	if err != nil {
		return types.ProductListingResponse{Offset: 0, Limit: 0, Result: []types.ProductResponse{}}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	var products []types.ProductResponse
	for rows.Next() {
		var product types.ProductResponse
		// Scan each row into the User struct
		err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Price, &product.Stock, &product.Description)
		if err != nil {
			return types.ProductListingResponse{Offset: 0, Limit: 0, Result: []types.ProductResponse{}}, fmt.Errorf("failed to scan row: %w", err)
		}
		products = append(products, product)
	}
	// Check for errors after the loop
	if err := rows.Err(); err != nil {
		return types.ProductListingResponse{Offset: 0, Limit: 0, Result: []types.ProductResponse{}}, fmt.Errorf("error iterating rows: %w", err)
	}
	return types.ProductListingResponse{Offset: limit, Limit: offset,
		Result: products, OrderBy: orderBy, SortBy: sortBy}, nil
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

func GetProductById(Id string) (types.Product, error) {
	pool := GetDBPool()

	query := `SELECT NAME, DESCRIPTION, PRICE, STOCK, CATEGORY FROM PRODUCTS WHERE ID=$1`
	var product types.Product
	err := pool.QueryRow(context.Background(), query, Id).Scan(&product.Name, &product.Description, &product.Price, &product.Stock, &product.Category)

	if err != nil {
		return types.Product{}, err
	}

	product.Id = Id
	return product, nil
}
