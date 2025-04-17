package v1

import (
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/repository/item"
	"github.com/webbsalad/pvz/internal/repository/pvz"
	pvz_service "github.com/webbsalad/pvz/internal/service/pvz"
)

type Service struct {
	pvzRepository  pvz.Repository
	itemRepository item.Repository

	config config.Config
}

func NewService(
	pvzRepository pvz.Repository,
	itemRepository item.Repository,
	config config.Config) pvz_service.Service {
	return &Service{
		pvzRepository:  pvzRepository,
		itemRepository: itemRepository,
		config:         config,
	}
}
