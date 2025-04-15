package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) CreatePVZ(ctx context.Context, pvz model.PVZ) (model.PVZ, error) {
	newPVZ, err := s.pvzRepository.CreatePVZ(ctx, pvz)
	if err != nil {
		return model.PVZ{}, fmt.Errorf("create pvz: %w", err)
	}

	return newPVZ, nil
}
