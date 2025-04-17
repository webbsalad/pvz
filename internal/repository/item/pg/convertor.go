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

	pvzID, err := model.NewPVZID(in.PVZID)
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

func toProductsFromDB(in []Product) ([]model.Product, error) {
	products := make([]model.Product, len(in))
	for i, dbProduct := range in {
		product, err := toProductFromDB(dbProduct)
		if err != nil {
			return nil, fmt.Errorf("convert db product to model: %w", err)
		}

		products[i] = product
	}

	return products, nil
}

func toProductFromDB(in Product) (model.Product, error) {
	productID, err := model.NewProductID(in.ID)
	if err != nil {
		return model.Product{}, fmt.Errorf("convert str product id to model: %w", err)
	}

	receptionID, err := model.NewReceptionID(in.ReceptionID)
	if err != nil {
		return model.Product{}, fmt.Errorf("convert str reception id to model: %w", err)
	}

	product := model.Product{
		ID:          productID,
		ReceptionID: receptionID,
		Type:        in.Type,

		DateTime: in.DateTime,
	}

	return product, nil
}
