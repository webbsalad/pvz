package user

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User, passhash string) (model.User, error)
}
