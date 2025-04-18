package pvz

import (
	"context"

	"github.com/webbsalad/pvz/internal/convertor"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPVZList(ctx context.Context, req *desc.GetPVZListRequest) (*desc.GetPVZListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	pvzs, err := i.pvzService.GetPVZList(ctx)
	if err != nil {
		return nil, convertor.ConvertError(err, i.log)
	}

	return &desc.GetPVZListResponse{
		Pvzs: convertor.ToDescsFromPVZs(pvzs),
	}, nil
}
