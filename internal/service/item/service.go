package item

import (
	"context"

	"github.com/webbsalad/pvz/internal/model"
)

type Service interface {
	CreateReception(ctx context.Context, role model.Role, pvzID model.PVZID) (model.Reception, error)
	AddProduct(ctx context.Context, role model.Role, pvzID model.PVZID, productType string) (model.Product, error)
	RemoveProduct(ctx context.Context, role model.Role, pvzID model.PVZID) error
	CloseReception(ctx context.Context, userRole model.Role, pvzID model.PVZID) (model.Reception, error)
}
