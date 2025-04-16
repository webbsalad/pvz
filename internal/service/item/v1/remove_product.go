package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) RemoveProduct(ctx context.Context, role model.Role, pvzID model.PVZID) error {
	if role != model.EMPLOYEE {
		return model.ErrWrongRole
	}

	reception, err := s.itemReporitory.GetReceptionsByParams(ctx, model.Reception{
		PVZID:  pvzID,
		Status: model.IN_PROGRESS,
	})
	if err != nil {
		return fmt.Errorf("get reception: %w", err)
	}

	if err := s.itemReporitory.RemoveProduct(ctx, reception[0].ID); err != nil {
		return fmt.Errorf("remove product: %w", err)
	}

	return nil
}
