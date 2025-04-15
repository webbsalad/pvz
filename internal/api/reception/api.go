package reception

import (
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/reception"
)

type Implementation struct {
	desc.UnimplementedReceptionServiceServer

	receptionService reception.Service
}

func NewImplementation(receptionService reception.Service) desc.ReceptionServiceServer {
	return &Implementation{
		receptionService: receptionService,
	}
}
