package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) CreateReception(ctx context.Context, role model.Role, pvzID model.PVZID) (model.Reception, error) {
	if role != model.EMPLOYEE {
		return model.Reception{}, model.ErrWrongRole
	}

	_, err := s.itemReporitory.GetReceptionsByParams(ctx, model.Reception{
		PVZID:  pvzID,
		Status: model.IN_PROGRESS,
	})
	if err != nil {
		if !errors.Is(err, model.ErrReceptionNotFound) {
			return model.Reception{}, fmt.Errorf("get in progress receptions: %w", err)
		}
	} else {
		return model.Reception{}, fmt.Errorf("reception with status in_progress now exist")
	}

	reception, err := s.itemReporitory.CreateReception(ctx, pvzID)
	if err != nil {
		return model.Reception{}, fmt.Errorf("create reception: %w", err)
	}

	return reception, nil
}
