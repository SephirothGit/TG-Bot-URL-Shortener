package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerAddr  string
}

func Load() (*Config, error) {

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("url is required")
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	return &Config{DatabaseURL: dbURL, ServerAddr: serverAddr}, nil
}
