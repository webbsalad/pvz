package login

import (
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/service/login"
)

type Implementation struct {
	desc.UnimplementedLoginServiceServer

	loginService login.Service
}

func NewImplementation(loginService login.Service) desc.LoginServiceServer {
	return &Implementation{
		loginService: loginService,
	}
}
