package v1

import (
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/repository/item"
	"github.com/webbsalad/pvz/internal/repository/pvz"
	"github.com/webbsalad/pvz/internal/repository/user"
	pvz_service "github.com/webbsalad/pvz/internal/service/pvz"
)

type Service struct {
	pvzRepository  pvz.Repository
	userRepository user.Repository
	itemRepository item.Repository

	config config.Config
}

func NewService(
	pvzRepository pvz.Repository,
	userRepository user.Repository,
	itemRepository item.Repository,
	config config.Config) pvz_service.Service {
	return &Service{
		pvzRepository:  pvzRepository,
		userRepository: userRepository,
		itemRepository: itemRepository,
		config:         config,
	}
}
