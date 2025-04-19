package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/webbsalad/pvz/internal/metrics"
	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) CreateReception(ctx context.Context, userRole model.Role, pvzID model.PVZID) (model.Reception, error) {
	if userRole != model.EMPLOYEE {
		return model.Reception{}, model.ErrWrongRole
	}

	status := model.IN_PROGRESS
	_, err := s.itemRepository.GetReceptionsByParams(ctx, model.ReceptionFilter{
		PVZID:  &pvzID,
		Status: &status,
	})
	if err != nil {
		if !errors.Is(err, model.ErrReceptionNotFound) {
			return model.Reception{}, fmt.Errorf("get in progress receptions: %w", err)
		}
	} else {
		return model.Reception{}, model.ErrReceptionAlreadyExist
	}

	reception, err := s.itemRepository.CreateReception(ctx, pvzID)
	if err != nil {
		return model.Reception{}, fmt.Errorf("create reception: %w", err)
	}

	metrics.ReceptionCreated.Inc()
	return reception, nil
}
