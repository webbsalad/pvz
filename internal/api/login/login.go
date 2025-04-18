package login

import (
	"context"

	"github.com/webbsalad/pvz/internal/convertor"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	token, err := i.loginService.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	return &desc.LoginResponse{
		Token: token,
	}, nil

}
