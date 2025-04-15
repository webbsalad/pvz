package app

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/webbsalad/pvz/internal/config"
)

func initDB(cfg config.Config) *sqlx.DB {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		log.Fatalf("failed connect to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to verify connection to db: %v", err)
	}

	return db
}
