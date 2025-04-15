package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/pvz/internal/repository/user"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (user.Repository, error) {
	return &Repository{db: db}, nil
}
