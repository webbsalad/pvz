package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/pvz/internal/repository/login"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (login.Repository, error) {
	return &Repository{db: db}, nil
}
