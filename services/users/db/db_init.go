package db

import (
	"context"
	"fmt"

	"go-microservices/users/config"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitializeDb(cfg config.Config) {

	connString := "postgres://" + cfg.Db_User + ":" + cfg.Db_Pwd + "@" + cfg.Db_URL + ":" + cfg.Db_Port + "/admin?sslmode=disable"
	var err error
	ctx := context.Background()
	dbPool, err = pgxpool.New(ctx, connString)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	/* defer dbPool.Close() */
	filePath := getSchemaFilePath()
	if err = executeSQLFromFile(ctx, dbPool, filePath); err != nil {
		log.Fatalf("Failed to execute SQL file %v", err)
	}
	log.Println("SQL file executed successfully.")
}

func getSchemaFilePath() string {
	// Check for an environment variable first
	if path := os.Getenv("SCHEMA_FILE_PATH"); path != "" {
		return path
	}

	// Fallback to executable-based path
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	baseDir := filepath.Dir(exePath)
	filePath := filepath.Join(baseDir, "internal", "db", "schema.sql")

	// Validate the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("Schema file not found at path: %s", filePath)
	}
	return filePath
}

func executeSQLFromFile(ctx context.Context, pool *pgxpool.Pool, filePath string) error {
	// Read the SQL file
	sqlBytes, err := os.ReadFile(filePath)

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

func GetDBPool() *pgxpool.Pool {
	return dbPool
}
