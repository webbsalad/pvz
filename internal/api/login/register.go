package login

import (
	"context"
	"net/http"

	"github.com/webbsalad/pvz/internal/convertor"
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/utils/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	role, err := model.NewRole(req.GetRole())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	user, err := i.loginService.Register(ctx, model.User{Email: req.GetEmail(), Role: role}, req.GetPassword())
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	if err := metadata.SetHTTPStatus(ctx, http.StatusCreated); err != nil {
		return nil, status.Errorf(codes.Internal, "set http status: %v", err)
	}

	return &desc.RegisterResponse{
		Id:    user.ID.String(),
		Email: user.Email,
		Role:  user.Role.String(),
	}, nil

}
