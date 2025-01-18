package config

import (
	"log"
	"os"
)

var cfg Config

type Config struct {
	Port    string
	Db_User string
	Db_Pwd  string
	Db_Port string
	Db_URL  string
}

func Load() Config {
	// Get the current environment (default to "development" if not set)
	env := getEnv("ENV", "development")

	// Load configuration based on the environment
	switch env {
	case "development":
		cfg = Config{
			Port:    getEnv("USER_SERVICE_PORT", "8080"),
			Db_User: getEnv("USER_SERVICE_DB_USER", "admin"),
			Db_Pwd:  getEnv("USER_SERVICE_DB_PWD", "admin"),
			Db_Port: getEnv("USER_SERVICE_DB_PORT", "5432"),
			Db_URL:  getEnv("USER_SERVICE_DB_HOST", "localhost"),
		}
	case "production":
		cfg = Config{
			Port:    getEnv("USER_SERVICE_PORT", "80"),
			Db_User: getEnv("USER_SERVICE_DB_USER", "admin"),
			Db_Pwd:  getEnv("USER_SERVICE_DB_PWD", "admin"),
			Db_Port: getEnv("USER_SERVICE_DB_PORT", "5432"),
			Db_URL:  getEnv("USER_SERVICE_DB_HOST", "localhost"),
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
