package pvz

import (
	"context"
	"time"

	"github.com/webbsalad/pvz/internal/model"
)

type Service interface {
	CreatePVZ(ctx context.Context, role model.Role, pvz model.PVZ) (model.PVZ, error)
	GetPVZList(ctx context.Context) ([]model.PVZ, error)
	GetPVZIntervalList(ctx context.Context, userRole model.Role, page, limit *int32, from, to *time.Time) ([]model.PVZWithReceptions, error)
}
