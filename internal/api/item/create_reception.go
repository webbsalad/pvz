package reception

import (
	"context"

	"github.com/webbsalad/pvz/internal/convertor"
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) CreateReception(ctx context.Context, req *desc.CreateReceptionRequest) (*desc.CreateReceptionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	pvzID, err := model.NewPVZID(req.GetPvzId())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	reception, err := i.itemService.CreateReception(ctx, pvzID)
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &desc.CreateReceptionResponse{
		Id:       reception.ID.String(),
		DateTime: timestamppb.New(reception.DateTime),
		PvzID:    reception.PVZID.String(),
		Status:   reception.Status.String(),
	}, nil
}
