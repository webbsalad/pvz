package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) GetPVZList(ctx context.Context) ([]model.PVZ, error) {
	pvzs, err := s.pvzRepository.GetPVZsByParams(ctx, model.PVZ{})
	if err != nil {
		return nil, fmt.Errorf("get pvzs: %w", err)
	}

	return pvzs, nil
}
