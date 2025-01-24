package config

import (
	"log"
	"os"
)

var cfg Config

type Config struct {
	Port        string
	Db_User     string
	Db_Pwd      string
	Db_Port     string
	Db_URL      string
	Redis_URL   string
	Schema_Path string
	Kafka_URL   string
}

func Load() Config {
	// Get the current environment (default to "development" if not set)
	env := getEnv("ENV", "development")
	// Load configuration based on the environment
	switch env {
	case "development":
		cfg = Config{
			Port:        getEnv("PRODUCT_SERVICE_PORT", "8081"),
			Db_User:     getEnv("PRODUCT_SERVICE_DB_USER", "admin"),
			Db_Pwd:      getEnv("PRODUCT_SERVICE_DB_PWD", "admin"),
			Db_Port:     getEnv("PRODUCT_SERVICE_DB_PORT", "5432"),
			Db_URL:      getEnv("PRODUCT_SERVICE_DB_HOST", "localhost"),
			Redis_URL:   getEnv("PRODUCT_SERVICE_REDIS_URL", "localhost:6379"),
			Schema_Path: getEnv("PRODUCT_SERVICE_SCHEMA_PATH", "C:\\Users\\vivek\\Documents\\go-microservices\\services\\product\\db\\schema.sql"),
			Kafka_URL:   getEnv("PRODUCT_SERVICE_KAFKA_URL", "localhost:9092"),
		}
	case "production":
		cfg = Config{
			Port:        getEnv("PRODUCT_SERVICE_PORT", "80"),
			Db_User:     getEnv("PRODUCT_SERVICE_DB_USER", "admin"),
			Db_Pwd:      getEnv("PRODUCT_SERVICE_DB_PWD", "admin"),
			Db_Port:     getEnv("PRODUCT_SERVICE_DB_PORT", "5432"),
			Db_URL:      getEnv("PRODUCT_SERVICE_DB_HOST", "localhost"),
			Redis_URL:   getEnv("PRODUCT_SERVICE_REDIS_URL", "localhost:6379"),
			Schema_Path: getEnv("PRODUCT_SERVICE_SCHEMA_PATH", "C:\\Users\\vivek\\Documents\\go-microservices\\services\\product\\db\\schema.sql"),
			Kafka_URL:   getEnv("PRODUCT_SERVICE_KAFKA_URL", "localhost:9092"),
		}
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	log.Printf("Configuration loaded for environment '%s': %+v", env, cfg)
	return cfg
}

// Helper function to get environment variables or return default values
func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func GetConfig() Config {
	return cfg
}
