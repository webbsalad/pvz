package reception

import (
	"log/slog"

	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/item"
)

type Implementation struct {
	desc.UnimplementedItemServiceServer

	itemService item.Service
	log         *slog.Logger
}

func NewImplementation(itemService item.Service, log *slog.Logger) desc.ItemServiceServer {
	return &Implementation{
		itemService: itemService,
		log:         log,
	}
}
