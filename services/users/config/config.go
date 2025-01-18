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
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	db_user := os.Getenv("USER_SERVICE_DB_USER")

	if db_user == "" {
		db_user = "admin" // Default user
	}

	db_pwd := os.Getenv("USER_SERVICE_DB_PWD")

	if db_pwd == "" {
		db_pwd = "admin" // Default password
	}

	db_port := os.Getenv("USER_SERVICE_DB_PORT")

	if db_port == "" {
		db_port = "5432" // Default password
	}

	db_url := os.Getenv("USER_SERVICE_DB_HOST")

	if db_url == "" {
		db_url = "localhost" // Default password
	}

	cfg = Config{Port: port, Db_User: db_user, Db_Pwd: db_pwd, Db_Port: db_port, Db_URL: db_url}
	log.Printf("Configuration loaded: %+v", cfg)
	return cfg
}

func GetConfig() Config {
	return cfg
}
