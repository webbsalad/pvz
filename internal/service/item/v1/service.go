package v1

import (
	"github.com/webbsalad/pvz/internal/config"
	"github.com/webbsalad/pvz/internal/repository/item"
	item_service "github.com/webbsalad/pvz/internal/service/item"
)

type Service struct {
	itemReporitory item.Repository

	config config.Config
}

func NewService(
	itemReporitory item.Repository,
	config config.Config) item_service.Service {
	return &Service{
		itemReporitory: itemReporitory,
		config:         config,
	}
}
