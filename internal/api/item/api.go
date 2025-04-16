package reception

import (
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/item"
)

type Implementation struct {
	desc.UnimplementedItemServiceServer

	itemService item.Service
}

func NewImplementation(itemService item.Service) desc.ItemServiceServer {
	return &Implementation{
		itemService: itemService,
	}
}
