package app

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/webbsalad/pvz/internal/config"
	"go.uber.org/fx"
)

func initDB(cfg config.Config, lc fx.Lifecycle) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifeMinutes) * time.Minute)

	log.Printf("DB pool configured: open=%d idle=%d life=%dmin",
		cfg.DBMaxOpenConns, cfg.DBMaxIdleConns, cfg.DBConnMaxLifeMinutes)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("closing DB")
			return db.Close()
		},
	})

	return db, nil
}
