package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/repository/pvz"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (pvz.Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) CreatePVZ(ctx context.Context, pvz model.PVZ) (model.PVZ, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("pvz").
		Columns("id", "city", "registration_date").
		Values(pvz.ID.String(), pvz.City, pvz.RegistrationDate).
		Suffix("RETURNING id, city, registration_date")

	q, args, err := query.ToSql()
	if err != nil {
		return model.PVZ{}, fmt.Errorf("build sql: %w", err)
	}

	var storedPVZ PVZ
	err = r.db.QueryRowContext(ctx, q, args...).Scan(&storedPVZ.ID, &storedPVZ.City, &storedPVZ.RegistrationDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.PVZ{}, model.ErrPVZAlreadyExist
		}
		return model.PVZ{}, fmt.Errorf("insert user: %w", err)
	}

	newPVZ, err := toPVZFromDB(storedPVZ)
	if err != nil {
		return model.PVZ{}, fmt.Errorf("convert stored pvz to model: %w", err)
	}

	return newPVZ, nil
}
