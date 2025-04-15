package login

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Service interface {
	DummyLogin(ctx context.Context, role model.Role) (string, error)
}
