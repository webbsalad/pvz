package convertor

import (
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDescFromReception(in model.Reception) *desc.Reception {
	return &desc.Reception{
		Id:       in.ID.String(),
		DateTime: timestamppb.New(in.DateTime),
		PvzID:    in.PVZID.String(),
		Status:   in.Status.String(),
	}
}
