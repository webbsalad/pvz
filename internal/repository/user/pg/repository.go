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

func (r *Repository) GetUserID(ctx context.Context, email string) (model.UserID, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Select("id").
		From("users").
		Where(
			sq.Eq{"email": email},
		)

	q, args, err := query.ToSql()
	if err != nil {
		return model.UserID{}, fmt.Errorf("build sql: %w", err)
	}

	var strUserID string
	if err := r.db.GetContext(ctx, &strUserID, q, args...); err != nil {
		return model.UserID{}, fmt.Errorf("get user id: %w", err)
	}

	userID, err := model.NewUserID(strUserID)
	if err != nil {
		return model.UserID{}, fmt.Errorf("convert stored user id to model: %w", err)
	}

	return userID, nil
}

func (r *Repository) GetPassHash(ctx context.Context, userID model.UserID) (string, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Select("passhash").
		From("users").
		Where(
			sq.Eq{"id": userID.String()},
		)

	q, args, err := query.ToSql()
	if err != nil {
		return "", fmt.Errorf("build sql: %w", err)
	}

	var passhash string
	if err := r.db.GetContext(ctx, &passhash, q, args...); err != nil {
		return "", fmt.Errorf("get passhash: %w", err)
	}

	return passhash, nil
}

func (r *Repository) GetUser(ctx context.Context, userID model.UserID) (model.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Select("*").
		From("users").
		Where(
			sq.Eq{"id": userID.String()},
		)

	q, args, err := query.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("build sql: %w", err)
	}

	var storedUser User
	if err := r.db.GetContext(ctx, &storedUser, q, args...); err != nil {
		return model.User{}, fmt.Errorf("get user: %w", err)
	}

	user, err := toUserFromDB(storedUser)
	if err != nil {
		return model.User{}, fmt.Errorf("convert str user id to model: %w", err)
	}

	return user, nil
}
