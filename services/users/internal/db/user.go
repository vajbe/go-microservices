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
	fmt.Print("Received pool")
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
