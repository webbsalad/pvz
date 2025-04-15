package pg

import (
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

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
