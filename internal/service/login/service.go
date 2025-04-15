package login

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Service interface {
	DummyLogin(ctx context.Context, role model.Role) (string, error)
	Register(ctx context.Context, user model.User, password string) (model.User, error)
}
