package pvz

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Service interface {
	CreatePVZ(ctx context.Context, role model.Role, pvz model.PVZ) (model.PVZ, error)
}
