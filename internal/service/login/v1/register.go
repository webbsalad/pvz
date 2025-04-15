package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
	"github.com/webbsalad/pvz/internal/utils/hash"
)

func (s *Service) Register(ctx context.Context, user model.User, password string) (model.User, error) {
	passhash, err := hash.HashPassword(password)
	if err != nil {
		return model.User{}, fmt.Errorf("hash password: %w", err)
	}

	newUser, err := s.userRepository.CreateUser(ctx, user, passhash)
	if err != nil {
		return model.User{}, fmt.Errorf("create user: %w", err)
	}

	return newUser, nil
}
