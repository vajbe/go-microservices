package db

import (
	"context"
	"fmt"
	"go-microservices/users/internal/types"
	"time"

	"log"
)

func AddUser(user types.User) {
	pool := GetDBPool()
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	var id int
	var createdAt time.Time

	err := pool.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&id, &createdAt)
	if err != nil {
		log.Printf("Error while inserting a record %s", err.Error())
		/* http.Error(w, "Failed to insert record", http.StatusInternalServerError)
		return */
	}
}

func GetUsers() ([]types.User, error) {
	pool := GetDBPool()

	rows, err := pool.Query(context.Background(), "SELECT * FROM users") // Replace "your_table_name" with your actual table name
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	var users []types.User
	for rows.Next() {
		var user types.User
		var createdAt time.Time
		var id int
		// Scan each row into the User struct
		err := rows.Scan(&id, &user.Name, &user.Email, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		user.CreatedAt = createdAt
		users = append(users, user)
	}
	// Check for errors after the loop
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	return users, nil
}
