package convertor

import (
	"github.com/webbsalad/pvz/internal/model"
	desc "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
)

func ToDescsFromPVZsWithReceptions(in []model.PVZWithReceptions) []*desc.PVZWithReceptions {
	out := make([]*desc.PVZWithReceptions, len(in))
	for i, v := range in {
		out[i] = ToDescFromPVZWithReceptions(v)
	}
	return out
}

func ToDescFromPVZWithReceptions(in model.PVZWithReceptions) *desc.PVZWithReceptions {
	return &desc.PVZWithReceptions{
		Pvz:        ToDescFromPVZ(in.PVZ),
		Receptions: ToDescsFromReceptionWithProducts(in.Receptions),
	}
}

func ToDescsFromReceptionWithProducts(in []model.ReceptionWithProducts) []*desc.ReceptionWithProducts {
	out := make([]*desc.ReceptionWithProducts, len(in))
	for i, v := range in {
		out[i] = ToDescFromReceptionWithProducts(v)
	}
	return out
}

func ToDescFromReceptionWithProducts(in model.ReceptionWithProducts) *desc.ReceptionWithProducts {
	return &desc.ReceptionWithProducts{
		Reception: ToDescFromReception(in.Reception),
		Products:  ToDescsFromProducts(in.Products),
	}
}
