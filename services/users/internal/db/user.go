package db

import (
	"context"
	"fmt"
	"go-microservices/users/internal/types"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func AddUser(user types.User) (types.User, error) {
	pool := GetDBPool()
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	var id int
	var createdAt time.Time

	err := pool.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&id, &createdAt)
	if err != nil {
		return user, fmt.Errorf("failed to insert record: %w", err)
	}
	return user, nil
}

func GetUsers() ([]types.User, error) {
	pool := GetDBPool()
	rows, err := pool.Query(context.Background(), "SELECT * FROM users")
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

func UpdateUser(user types.User) error {
	pool := GetDBPool()
	query := `UPDATE users set name=$1, email=$2 where id=$3`
	tag, err := pool.Exec(context.Background(), query, user.Name, user.Email, user.Id)
	if err != nil {
		return fmt.Errorf("error update record: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no record found with id %d", user.Id)
	}

	log.Print("User updated successfuly.")

	return nil
}

func GetUser(user types.User) (types.User, error) {
	pool := GetDBPool()
	query := "SELECT * FROM users where id=$1"
	var resultUser types.User
	err := pool.QueryRow(context.Background(), query, user.Id).Scan(&resultUser.Id, &resultUser.Name, &resultUser.Email, &resultUser.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return resultUser, fmt.Errorf("no record found for id: %d", user.Id)
		}
		return resultUser, fmt.Errorf("failed to execute query: %w", err)
	}

	return resultUser, nil
}

func DeleteUser(user types.User) error {
	pool := GetDBPool()
	query := `DELETE FROM USERS where id=$1`
	tag, err := pool.Exec(context.Background(), query, user.Id)
	if err != nil {
		return fmt.Errorf("error DELETE record: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no record found with id %d", user.Id)
	}

	log.Print("User deleted successfuly.")

	return nil
}
