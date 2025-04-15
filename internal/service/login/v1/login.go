package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/utils/hash"
	"github.com/webbsalad/pvz/internal/utils/jwt"
)

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	userID, err := s.userRepository.GetUserID(ctx, email)
	if err != nil {
		return "", fmt.Errorf("get user id: %w", err)
	}

	passhash, err := s.userRepository.GetPassHash(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get passhash: %w", err)
	}

	if err := hash.CheckPassword(passhash, password); err != nil {
		return "", fmt.Errorf("wrong password: %w", err)
	}

	user, err := s.userRepository.GetUser(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get user: %w", err)
	}

	token, err := jwt.GenerateTokens(user.Role, s.config.JWTSecret)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil

}
