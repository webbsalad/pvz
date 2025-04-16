package item

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Repository interface {
	GetReceptionsByParams(ctx context.Context, reception model.Reception) ([]model.Reception, error)
	CreateReception(ctx context.Context, pvzID model.PVZID) (model.Reception, error)
	AddProduct(ctx context.Context, product model.Product) (model.Product, error)
}
