package v1

import (
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/repository/item"
	"github.com/webbsalad/pvz/internal/repository/pvz"
	"github.com/webbsalad/pvz/internal/repository/user"
	item_service "github.com/webbsalad/pvz/internal/service/item"
)

type Service struct {
	pvzRepository  pvz.Repository
	itemReporitory item.Repository
	userRepository user.Repository

	config config.Config
}

func NewService(
	pvzRepository pvz.Repository,
	itemReporitory item.Repository,
	userRepository user.Repository,
	config config.Config) item_service.Service {
	return &Service{
		pvzRepository:  pvzRepository,
		itemReporitory: itemReporitory,
		userRepository: userRepository,
		config:         config,
	}
}
