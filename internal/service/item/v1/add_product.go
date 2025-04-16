package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) AddProduct(ctx context.Context, userRole model.Role, pvzID model.PVZID, productType string) (model.Product, error) {
	if userRole != model.EMPLOYEE {
		return model.Product{}, model.ErrWrongRole
	}

	receptions, err := s.itemReporitory.GetReceptionsByParams(ctx, model.Reception{
		PVZID:  pvzID,
		Status: model.IN_PROGRESS,
	})
	if err != nil {
		return model.Product{}, fmt.Errorf("get in progress reception: %w", err)
	}

	product, err := s.itemReporitory.AddProduct(ctx, model.Product{
		ReceptionID: receptions[0].ID,
		Type:        productType,
	})
	if err != nil {
		return model.Product{}, fmt.Errorf("add: %w", err)
	}

	return product, nil
}
