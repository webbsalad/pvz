package login

import (
	"context"

	"github.com/webbsalad/pvz/internal/convertor"
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DummyLogin(ctx context.Context, req *desc.DummyLoginRequest) (*desc.DummyLoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	role, err := model.NewRole(req.GetRole())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	token, err := i.loginService.DummyLogin(ctx, role)
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	return &desc.DummyLoginResponse{
		Token: token,
	}, nil
}
