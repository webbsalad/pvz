package v1

import (
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/repository/login"
	"github.com/webbsalad/pvz/internal/repository/user"
	login_service "github.com/webbsalad/pvz/internal/service/login"
)

type Service struct {
	loginRepository login.Repository
	userRepository  user.Repository

	config config.Config
}

func NewService(
	loginRepository login.Repository,
	userRepository user.Repository,
	config config.Config) login_service.Service {
	return &Service{
		loginRepository: loginRepository,
		userRepository:  userRepository,
		config:          config,
	}
}
