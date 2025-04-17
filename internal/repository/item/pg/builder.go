package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/webbsalad/pvz/internal/model"
)

func buildReceptionWhere(f model.ReceptionFilter) sq.And {
	where := sq.And{}

	if f.PVZID != nil {
		where = append(where, sq.Eq{"pvz_id": f.PVZID.String()})
	}
	if f.Status != nil {
		where = append(where, sq.Eq{"status": f.Status.String()})
	}
	if f.From != nil {
		where = append(where, sq.GtOrEq{"created_at": *f.From})
	}
	if f.To != nil {
		where = append(where, sq.LtOrEq{"created_at": *f.To})
	}

	return where
}

func buildProductWhere(f model.ProductFilter) sq.And {
	where := sq.And{}

	if f.ReceptionID != nil {
		where = append(where, sq.Eq{"reception_id": f.ReceptionID.String()})
	}
	if f.Type != nil {
		where = append(where, sq.Eq{"type": f.Type})
	}
	if f.From != nil {
		where = append(where, sq.GtOrEq{"created_at": *f.From})
	}
	if f.To != nil {
		where = append(where, sq.LtOrEq{"created_at": *f.To})
	}

	return where
}
