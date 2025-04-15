package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/utils/jwt"
)

func (s *Service) CreatePVZ(ctx context.Context, token string, pvz model.PVZ) (model.PVZ, error) {
	role, err := jwt.ExtractClaimsFromToken(token, s.config.JWTSecret)
	if err != nil {
		return model.PVZ{}, fmt.Errorf("get role from token: %w", err)
	}

	if role != model.MODERATOR {
		return model.PVZ{}, model.ErrWrongRole
	}

	newPVZ, err := s.pvzRepository.CreatePVZ(ctx, pvz)
	if err != nil {
		return model.PVZ{}, fmt.Errorf("create pvz: %w", err)
	}

	return newPVZ, nil
}
