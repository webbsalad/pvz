package pvz

import (
	"context"
	"net/http"

	"github.com/webbsalad/pvz/internal/convertor"
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/utils/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) CreatePVZ(ctx context.Context, req *desc.CreatePVZRequest) (*desc.CreatePVZResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	pvzID, err := model.NewPVZID(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	userRole, err := metadata.GetRole(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	pvz, err := i.pvzService.CreatePVZ(ctx, userRole, model.PVZ{
		ID:               pvzID,
		City:             req.GetCity(),
		RegistrationDate: req.GetRegistrationDate().AsTime(),
	})
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	if err := metadata.SetHTTPStatus(ctx, http.StatusCreated); err != nil {
		return nil, status.Errorf(codes.Internal, "set http status: %v", err)
	}

	return &desc.CreatePVZResponse{
		Id:               pvz.ID.String(),
		City:             pvz.City,
		RegistrationDate: timestamppb.New(pvz.RegistrationDate),
	}, nil
}
