package v1

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/webbsalad/pvz/internal/model"
)

func (s *Service) GetPVZIntervalList(ctx context.Context, userRole model.Role, page, limit *int32, from, to *time.Time) ([]model.PVZWithReceptions, error) {
	if userRole != model.EMPLOYEE && userRole != model.MODERATOR {
		return nil, model.ErrWrongRole
	}

	receptionFilter := model.ReceptionFilter{
		From: from,
		To:   to,
	}
	recs, err := s.itemRepository.GetReceptionsByParams(ctx, receptionFilter)
	if err != nil {
		if errors.Is(err, model.ErrReceptionNotFound) {
			return nil, model.ErrReceptionNotFound
		}
		return nil, fmt.Errorf("get receptions: %w", err)
	}

	pvzIDSet := make(map[model.PVZID]struct{}, len(recs))
	for _, r := range recs {
		pvzIDSet[r.PVZID] = struct{}{}
	}
	pvzIDs := make([]model.PVZID, 0, len(pvzIDSet))
	for id := range pvzIDSet {
		pvzIDs = append(pvzIDs, id)
	}

	pvzFilter := model.PVZFilter{
		IDs:   pvzIDs,
		Page:  page,
		Limit: limit,
	}
	pvzs, err := s.pvzRepository.GetPVZsByParams(ctx, pvzFilter)
	if err != nil {
		if errors.Is(err, model.ErrPVZNotFound) {
			return nil, model.ErrPVZNotFound
		}
		return nil, fmt.Errorf("get pvzs: %w", err)
	}

	recsWithProds := make([]model.ReceptionWithProducts, len(recs))
	for i, r := range recs {
		prods, err := s.itemRepository.GetProductssByParams(
			ctx,
			model.ProductFilter{ReceptionID: &r.ID},
		)

		if err != nil && !errors.Is(err, model.ErrProductNotFound) {
			return nil, fmt.Errorf("get products for reception %s: %w", r.ID, err)
		}

		recsWithProds[i] = model.ReceptionWithProducts{
			Reception: r,
			Products:  prods,
		}
	}

	recsByPVZ := make(map[model.PVZID][]model.ReceptionWithProducts, len(pvzIDs))
	for _, rwp := range recsWithProds {
		recsByPVZ[rwp.Reception.PVZID] = append(recsByPVZ[rwp.Reception.PVZID], rwp)
	}

	result := make([]model.PVZWithReceptions, len(pvzs))
	for i, pvz := range pvzs {
		result[i] = model.PVZWithReceptions{
			PVZ:        pvz,
			Receptions: recsByPVZ[pvz.ID],
		}
	}

	return result, nil
}
