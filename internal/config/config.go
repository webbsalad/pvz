package config

import (
	"log"
	"os"
)

type Config struct {
	DSN       string `yaml:"dsn"`
	JWTSecret string `yaml:"jwt_secret"`
}

func NewConfig() Config {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("environment variable DSN is required")
	}

	jwt := os.Getenv("JWT_SECRET")
	if jwt == "" {
		log.Fatal("environment variable JWT_SECRET is required")
	}

	return Config{
		DSN:       dsn,
		JWTSecret: jwt,
	}
}
