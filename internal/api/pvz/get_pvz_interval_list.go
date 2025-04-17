package pvz

import (
	"context"
	"time"

	"github.com/webbsalad/pvz/internal/convertor"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"github.com/webbsalad/pvz/internal/utils/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPVZIntervalList(ctx context.Context, req *desc.GetPVZIntervalListRequest) (*desc.GetPVZIntervalListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userRole, err := metadata.GetRole(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%v", err)
	}

	var (
		from  *time.Time
		to    *time.Time
		page  *int32
		limit *int32
	)

	if ts := req.GetStartDate(); ts != nil {
		t := ts.AsTime()
		from = &t
	}
	if ts := req.GetEndDate(); ts != nil {
		t := ts.AsTime()
		to = &t
	}
	if req.Page != nil {
		p := req.GetPage()
		page = &p
	}
	if req.Limit != nil {
		l := req.GetLimit()
		limit = &l
	}

	pvzs, err := i.pvzService.GetPVZIntervalList(ctx, userRole, page, limit, from, to)
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &desc.GetPVZIntervalListResponse{
		Pvzs: convertor.ToDescsFromPVZsWithReceptions(pvzs),
	}, nil
}
