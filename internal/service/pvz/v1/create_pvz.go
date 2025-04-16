package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) CreatePVZ(ctx context.Context, role model.Role, pvz model.PVZ) (model.PVZ, error) {
	if role != model.MODERATOR {
		return model.PVZ{}, model.ErrWrongRole
	}

	newPVZ, err := s.pvzRepository.CreatePVZ(ctx, pvz)
	if err != nil {
		return model.PVZ{}, fmt.Errorf("create pvz: %w", err)
	}

	return newPVZ, nil
}
