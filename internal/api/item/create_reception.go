package reception

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

func (i *Implementation) CreateReception(ctx context.Context, req *desc.CreateReceptionRequest) (*desc.CreateReceptionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userRole, err := metadata.GetRole(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	pvzID, err := model.NewPVZID(req.GetPvzId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	reception, err := i.itemService.CreateReception(ctx, userRole, pvzID)
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	if err := metadata.SetHTTPStatus(ctx, http.StatusCreated); err != nil {
		return nil, status.Errorf(codes.Internal, "set http status: %v", err)
	}
	return &desc.CreateReceptionResponse{
		Id:       reception.ID.String(),
		DateTime: timestamppb.New(reception.DateTime),
		PvzID:    reception.PVZID.String(),
		Status:   reception.Status.String(),
	}, nil
}
