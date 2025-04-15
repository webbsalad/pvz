package pvz

import (
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/pvz"
)

type Implementation struct {
	desc.UnimplementedPVZServiceServer

	pvzService pvz.Service
}

func NewImplementation(pvzService pvz.Service) desc.PVZServiceServer {
	return &Implementation{
		pvzService: pvzService,
	}
}
