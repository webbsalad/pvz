package reception

import (
	"context"

	"github.com/webbsalad/pvz/internal/convertor"
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/utils/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) RemoveProduct(ctx context.Context, req *desc.RemoveProductRequest) (*desc.RemoveProductResponse, error) {
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

	if err := i.itemService.RemoveProduct(ctx, userRole, pvzID); err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	return &desc.RemoveProductResponse{}, nil
}
