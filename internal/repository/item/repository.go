package item

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Repository interface {
	GetReceptionsByParams(ctx context.Context, receptionFilter model.ReceptionFilter) ([]model.Reception, error)
	GetProductsByParams(ctx context.Context, productFilter model.ProductFilter) ([]model.Product, error)
	CreateReception(ctx context.Context, pvzID model.PVZID) (model.Reception, error)
	AddProduct(ctx context.Context, product model.Product) (model.Product, error)
	RemoveProduct(ctx context.Context, receptionID model.ReceptionID) error
	UpdateReception(ctx context.Context, reception model.Reception) (model.Reception, error)
}
