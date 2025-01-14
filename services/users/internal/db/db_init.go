package db

import (
	"context"
	"fmt"
	"go-microservices/users/internal/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeDb(cfg config.Config) {

	connString := "postgres://" + cfg.Db_User + ":" + cfg.Db_Pwd + "@" + cfg.Db_URL + ":" + cfg.Db_Port + "/admin?sslmode=disable"

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connString)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer pool.Close()

	exePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	baseDir := filepath.Dir(exePath)
	filePath := filepath.Join(baseDir, "users", "internal", "db", "schema.sql")
	if err = executeSQLFromFile(ctx, pool, filePath); err != nil {
		log.Fatalf("Failed to execute SQL file %v", err)
	}
	log.Println("SQL file executed successfully.")
}

func executeSQLFromFile(ctx context.Context, pool *pgxpool.Pool, filePath string) error {
	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	sqlQuery := string(sqlBytes)
	// Execute the SQL
	_, err = pool.Exec(ctx, sqlQuery)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	log.Printf("Executed SQL file: %s", filePath)
	return nil
}
