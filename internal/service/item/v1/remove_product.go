package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) RemoveProduct(ctx context.Context, userRole model.Role, pvzID model.PVZID) error {
	if userRole != model.EMPLOYEE {
		return model.ErrWrongRole
	}

	status := model.IN_PROGRESS
	reception, err := s.itemRepository.GetReceptionsByParams(ctx, model.ReceptionFilter{
		PVZID:  &pvzID,
		Status: &status,
	})
	if err != nil {
		return fmt.Errorf("get reception: %w", err)
	}

	if err := s.itemRepository.RemoveProduct(ctx, reception[0].ID); err != nil {
		return fmt.Errorf("remove product: %w", err)
	}

	return nil
}
