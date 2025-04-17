package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) GetPVZIntervalList(ctx context.Context, userRole model.Role, filter model.PVZFilter) ([]model.PVZWithReceptions, error) {
	if userRole != model.EMPLOYEE && userRole != model.MODERATOR {
		return nil, model.ErrWrongRole
	}

	pvzs, err := s.pvzRepository.GetPVZsByParams(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get pvzs: %w", err)
	}

	pvzsWithRecs := make([]model.PVZWithReceptions, len(pvzs))
	for i, pvz := range pvzs {
		recs, err := s.itemRepository.GetReceptionsByParams(ctx, model.ReceptionFilter{
			PVZID: &pvz.ID,
		})
		if err != nil && !errors.Is(err, model.ErrReceptionNotFound) {
			return nil, fmt.Errorf("get recs: %w", err)
		}

		recsWithProd := make([]model.ReceptionWithProducts, len(recs))
		for j, rec := range recs {
			prods, err := s.itemRepository.GetProductssByParams(ctx, model.ProductFilter{
				ReceptionID: &rec.ID,
			})
			if err != nil && !errors.Is(err, model.ErrProductNotFound) {
				return nil, fmt.Errorf("get products: %w", err)
			}
			recsWithProd[j] = model.ReceptionWithProducts{
				Reception: rec,
				Products:  prods,
			}
		}

		pvzsWithRecs[i] = model.PVZWithReceptions{
			PVZ:        pvz,
			Receptions: recsWithProd,
		}
	}

	return pvzsWithRecs, nil
}
