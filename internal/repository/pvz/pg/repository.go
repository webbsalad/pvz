package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/pvz/internal/repository/pvz"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (pvz.Repository, error) {
	return &Repository{db: db}, nil
}
