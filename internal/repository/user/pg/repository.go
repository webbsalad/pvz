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
	"github.com/webbsalad/pvz/internal/repository/user"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (user.Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) CreateUser(ctx context.Context, user model.User, passhash string) (model.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("users").
		Columns("email", "passhash", "role", "created_at").
		Values(user.Email, passhash, user.Role.String(), time.Now()).
		Suffix("ON CONFLICT (email) DO NOTHING RETURNING id, email, role")

	q, args, err := query.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("build sql: %w", err)
	}

	var storedUser User
	err = r.db.QueryRowContext(ctx, q, args...).
		Scan(&storedUser.ID, &storedUser.Email, &storedUser.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, model.ErrUserAlreadyExist
		}
		return model.User{}, fmt.Errorf("insert user: %w", err)
	}

	newStoredUser, err := toUserFromDB(storedUser)
	if err != nil {
		return model.User{}, fmt.Errorf("convert stored user to model: %w", err)
	}

	return newStoredUser, nil
}
