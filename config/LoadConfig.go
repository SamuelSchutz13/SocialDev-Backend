package configs

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required but not set")
	}

	return &Config{
		JWTSecret: jwtSecret,
	}
}
