package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/pvz/internal/repository/reception"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (reception.Repository, error) {
	return &Repository{db: db}, nil
}
