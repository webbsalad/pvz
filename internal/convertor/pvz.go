package convertor

import (
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDescsFromPVZs(in []model.PVZ) []*desc.PVZ {
	pvzs := make([]*desc.PVZ, len(in))
	for i, pvz := range in {
		pvzs[i] = ToDescFromPVZ(pvz)
	}
	return pvzs
}

func ToDescFromPVZ(in model.PVZ) *desc.PVZ {
	return &desc.PVZ{
		Id:               in.ID.String(),
		RegistrationDate: timestamppb.New(in.RegistrationDate),
		City:             in.City,
	}
}
