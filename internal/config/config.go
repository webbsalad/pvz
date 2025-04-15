package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string `yaml:"dsn"`
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	return Config{
		DSN: os.Getenv("DSN"),
	}
}
