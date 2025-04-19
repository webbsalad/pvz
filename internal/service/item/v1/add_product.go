package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/metrics"
	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) AddProduct(ctx context.Context, userRole model.Role, pvzID model.PVZID, product model.Product) (model.Product, error) {
	if userRole != model.EMPLOYEE {
		return model.Product{}, model.ErrWrongRole
	}

	status := model.IN_PROGRESS
	receptions, err := s.itemRepository.GetReceptionsByParams(ctx, model.ReceptionFilter{
		PVZID:  &pvzID,
		Status: &status,
	})
	if err != nil {
		return model.Product{}, fmt.Errorf("get in progress reception: %w", err)
	}

	product.ReceptionID = receptions[0].ID
	newProduct, err := s.itemRepository.AddProduct(ctx, product)
	if err != nil {
		return model.Product{}, fmt.Errorf("add: %w", err)
	}

	metrics.ProductAdded.Inc()
	return newProduct, nil
}
