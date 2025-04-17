package convertor

import (
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDescFromProduct(in model.Product) *desc.Product {
	return &desc.Product{
		Id:          in.ID.String(),
		DateTime:    timestamppb.New(in.DateTime),
		Type:        in.Type,
		ReceptionId: in.ReceptionID.String(),
	}
}

func ToDescsFromProducts(in []model.Product) []*desc.Product {
	out := make([]*desc.Product, len(in))
	for i, v := range in {
		out[i] = ToDescFromProduct(v)
	}
	return out
}
