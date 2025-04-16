package pg

import (
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func toPVZsFromDB(in []PVZ) ([]model.PVZ, error) {
	pvzs := make([]model.PVZ, len(in))
	for i, dbPVZ := range in {
		pvz, err := toPVZFromDB(dbPVZ)
		if err != nil {
			return nil, fmt.Errorf("convert db pvz to model: %w", err)
		}

		pvzs[i] = pvz
	}

	return pvzs, nil
}

func toPVZFromDB(in PVZ) (model.PVZ, error) {
	pvzID, err := model.NewPVZID(in.ID)
	if err != nil {
		return model.PVZ{}, fmt.Errorf("convert str to pvz id: %w", err)
	}

	pvz := model.PVZ{
		ID:   pvzID,
		City: in.City,

		RegistrationDate: in.RegistrationDate,
	}

	return pvz, nil
}
