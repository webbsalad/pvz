package pvz

import (
	"log/slog"

	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/pvz"
)

type Implementation struct {
	desc.UnimplementedPVZServiceServer

	pvzService pvz.Service
	log        *slog.Logger
}

func NewImplementation(pvzService pvz.Service, log *slog.Logger) desc.PVZServiceServer {
	return &Implementation{
		pvzService: pvzService,
		log:        log,
	}
}
