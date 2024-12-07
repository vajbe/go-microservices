package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func Load() Config {
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Configuration loaded: PORT=%s", port)
	return Config{Port: port}
}
