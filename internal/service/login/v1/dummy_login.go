package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/utils/jwt"
)

func (s *Service) DummyLogin(ctx context.Context, role model.Role) (string, error) {
	token, err := jwt.GenerateTokens(role, s.config.JWTSecret)
	if err != nil {
		return "", fmt.Errorf("generate tokens: %w", err)
	}

	return token, nil
}
