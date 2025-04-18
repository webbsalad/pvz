// internal/config/config.go
package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DSN                  string `yaml:"dsn"`
	JWTSecret            string `yaml:"jwt_secret"`
	DBMaxOpenConns       int    `yaml:"db_max_open_conns"`
	DBMaxIdleConns       int    `yaml:"db_max_idle_conns"`
	DBConnMaxLifeMinutes int    `yaml:"db_conn_max_life_minutes"`
}

func mustIntEnv(name string, defaultVal int) int {
	s := os.Getenv(name)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("invalid %s=%q: %v", name, s, err)
	}
	return v
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
		DSN:                  dsn,
		JWTSecret:            jwt,
		DBMaxOpenConns:       mustIntEnv("DB_MAX_OPEN_CONNS", 100),
		DBMaxIdleConns:       mustIntEnv("DB_MAX_IDLE_CONNS", 50),
		DBConnMaxLifeMinutes: mustIntEnv("DB_CONN_MAX_LIFE_MINUTES", 30),
	}
}
