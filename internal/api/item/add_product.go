package reception

import (
	"context"

	"github.com/webbsalad/pvz/internal/convertor"
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/utils/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) AddProduct(ctx context.Context, req *desc.AddProductRequest) (*desc.AddProductResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userRole, err := metadata.GetRole(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	pvzID, err := model.NewPVZID(req.GetPvzId())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	product, err := i.itemService.AddProduct(ctx, userRole, pvzID, req.GetType())
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	return &desc.AddProductResponse{
		Id:          product.ID.String(),
		DateTime:    timestamppb.New(product.DateTime),
		Type:        product.Type,
		ReceptionId: product.ReceptionID.String(),
	}, nil
}
