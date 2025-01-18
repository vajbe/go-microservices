package db

import (
	"context"
	"fmt"
	"go-microservices/users/types"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user types.User) (types.UserResponse, error) {
	pool := GetDBPool()
	query := `INSERT INTO users (name, email, password_hash, phone_number) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	var id string
	var createdAt int64
	/* Generate password hash here */

	validator := validator.New()
	err := validator.Struct(user)

	if err != nil {
		return types.UserResponse{}, fmt.Errorf("validation failed: %w", err)
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return types.UserResponse{Name: user.Name, Email: user.Email}, err
	}

	err = pool.QueryRow(context.Background(), query, user.Name, user.Email, string(password_hash), user.Phone).Scan(&id, &createdAt)
	if err != nil {
		return types.UserResponse{Name: user.Name, Email: user.Email}, fmt.Errorf("failed to insert record: %w", err)
	}
	response := types.UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Id:    id,
	}
	return response, nil
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
		var createdAt int64
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
		return fmt.Errorf("no record found with id %s", user.Id)
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
			return resultUser, fmt.Errorf("no record found for id: %s", user.Id)
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
		return fmt.Errorf("no record found with id %s", user.Id)
	}

	log.Print("User deleted successfuly.")

	return nil
}

func verifyUser(user types.UserLogin) (bool, error) {
	pool := GetDBPool()
	var password_hash string

	err := pool.QueryRow(context.Background(), `SELECT PASSWORD_HASH FROM USERS WHERE EMAIL=$1`, user.Name).Scan(&password_hash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, fmt.Errorf("invalid credentials")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(user.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("invalid credentials")
		}
		return false, err
	}

	return true, nil
}

func Login(user types.UserLogin) (types.UserLoginResponse, error) {
	validator := validator.New()
	err := validator.Struct(user)
	if err != nil {
		return types.UserLoginResponse{Name: user.Name}, err
	}
	isValid, err := verifyUser(user)
	if err != nil {
		return types.UserLoginResponse{Name: user.Name}, err
	}
	if !isValid {
		return types.UserLoginResponse{Name: user.Name}, fmt.Errorf("invalid credentials")
	}
	return types.UserLoginResponse{Name: user.Name}, nil
}
