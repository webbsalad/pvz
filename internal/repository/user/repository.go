package user

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User, passhash string) (model.User, error)
	GetUserID(ctx context.Context, email string) (model.UserID, error)
	GetUser(ctx context.Context, userID model.UserID) (model.User, error)
	GetPassHash(ctx context.Context, userID model.UserID) (string, error)
}
