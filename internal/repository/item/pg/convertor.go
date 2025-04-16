package pg

import (
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func toReceptionsFromDB(in []Reception) ([]model.Reception, error) {
	receptions := make([]model.Reception, len(in))
	for i, dbReception := range in {
		reception, err := toReceptionFromDB(dbReception)
		if err != nil {
			return nil, fmt.Errorf("convert db reception to model: %w", err)
		}

		receptions[i] = reception
	}

	return receptions, nil
}

func toReceptionFromDB(in Reception) (model.Reception, error) {
	receptionID, err := model.NewReceptionID(in.ID)
	if err != nil {
		return model.Reception{}, fmt.Errorf("convert str reception id to model: %w", err)
	}

	pvzID, err := model.NewPVZID(in.ID)
	if err != nil {
		return model.Reception{}, fmt.Errorf("convert str to pvz id: %w", err)
	}

	status, err := model.NewStatus(in.Status)
	if err != nil {
		return model.Reception{}, fmt.Errorf("convert str status to model: %w", err)
	}

	reception := model.Reception{
		ID:     receptionID,
		PVZID:  pvzID,
		Status: status,

		DateTime: in.DateTime,
	}

	return reception, nil
}
