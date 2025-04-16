package pvz

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Repository interface {
	CreatePVZ(ctx context.Context, pvz model.PVZ) (model.PVZ, error)
	GetPVZsByParams(ctx context.Context, pvz model.PVZ) ([]model.PVZ, error)
}
