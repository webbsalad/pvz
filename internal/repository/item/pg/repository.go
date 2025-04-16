package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/repository/item"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (item.Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) GetReceptionsByParams(ctx context.Context, reception model.Reception) ([]model.Reception, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	whereClause, err := buildReceptionWhere(reception)
	if err != nil {
		return nil, fmt.Errorf("build where clause: %w", err)
	}

	query := psql.
		Select("*").
		From("reception")

	if len(whereClause) > 0 {
		query = query.Where(whereClause)
	}

	q, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build sql: %w", err)
	}

	var storedReceptions []Reception
	if err := r.db.SelectContext(ctx, &storedReceptions, q, args...); err != nil {
		return nil, fmt.Errorf("select from receptions: %w", err)
	}

	if len(storedReceptions) == 0 {
		return nil, model.ErrReceptionNotFound
	}

	receptions, err := toReceptionsFromDB(storedReceptions)
	if err != nil {
		return nil, fmt.Errorf("convert stored receptions to model: %w", err)
	}

	return receptions, nil
}

func (r *Repository) CreateReception(ctx context.Context, pvzID model.PVZID) (model.Reception, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("reception").
		Columns("pvz_id", "status", "date_time").
		Values(pvzID.String(), model.IN_PROGRESS.String(), time.Now()).
		Suffix("ON CONFLICT (pvz_id) WHERE status = 'in_progress' DO NOTHING RETURNING id, pvz_id, status, date_time")

	q, args, err := query.ToSql()
	if err != nil {
		return model.Reception{}, fmt.Errorf("build sql: %w", err)
	}

	var storedReception Reception
	err = r.db.QueryRowContext(ctx, q, args...).
		Scan(&storedReception.ID, &storedReception.PVZID, &storedReception.Status, &storedReception.DateTime)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Reception{}, model.ErrReceptionAlreadyExist
		}
		return model.Reception{}, fmt.Errorf("insert reception: %w", err)
	}

	newStoredReception, err := toReceptionFromDB(storedReception)
	if err != nil {
		return model.Reception{}, fmt.Errorf("convert stored user to model: %w", err)
	}

	return newStoredReception, nil
}
