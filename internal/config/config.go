package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddr    string
	TelegramToken string
	JWTSecret     string
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

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("jwt secret is required")
	}

	return &Config{DatabaseURL: dbURL, ServerAddr: serverAddr, TelegramToken: telegramToken, JWTSecret: jwtSecret}, nil
}
