package v1

import (
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/repository/pvz"
	"github.com/webbsalad/pvz/internal/repository/reception"
	"github.com/webbsalad/pvz/internal/repository/user"
	reception_service "github.com/webbsalad/pvz/internal/service/reception"
)

type Service struct {
	pvzRepository       pvz.Repository
	receptionReporitory reception.Repository
	userRepository      user.Repository

	config config.Config
}

func NewService(
	pvzRepository pvz.Repository,
	receptionReporitory reception.Repository,
	userRepository user.Repository,
	config config.Config) reception_service.Service {
	return &Service{
		pvzRepository:       pvzRepository,
		receptionReporitory: receptionReporitory,
		userRepository:      userRepository,
		config:              config,
	}
}
