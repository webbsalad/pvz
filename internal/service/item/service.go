package item

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Service interface {
	CreateReception(ctx context.Context, pvzID model.PVZID) (model.Reception, error)
}
