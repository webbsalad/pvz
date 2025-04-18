package login

import (
	"log/slog"

	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/login"
)

type Implementation struct {
	desc.UnimplementedLoginServiceServer

	loginService login.Service
	log          *slog.Logger
}

func NewImplementation(loginService login.Service, log *slog.Logger) desc.LoginServiceServer {
	return &Implementation{
		loginService: loginService,
		log:          log,
	}
}
