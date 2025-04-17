package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) CloseReception(ctx context.Context, userRole model.Role, pvzID model.PVZID) (model.Reception, error) {
	if userRole != model.EMPLOYEE {
		return model.Reception{}, model.ErrWrongRole
	}

	status := model.IN_PROGRESS
	reception, err := s.itemRepository.GetReceptionsByParams(ctx, model.ReceptionFilter{
		PVZID:  &pvzID,
		Status: &status,
	})
	if err != nil {
		return model.Reception{}, fmt.Errorf("get receptions: %w", err)
	}

	newReception := reception[0]
	newReception.Status = model.CLOSE

	updatedReception, err := s.itemRepository.UpdateReception(ctx, newReception)
	if err != nil {
		return model.Reception{}, fmt.Errorf("close reception: %w", err)
	}

	return updatedReception, nil
}
