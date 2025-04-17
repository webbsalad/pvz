package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/webbsalad/pvz/internal/model"
)

func buildPVZWhere(f model.PVZFilter) sq.And {
	where := sq.And{}

	if f.City != nil {
		where = append(where, sq.Eq{"pvz_id": f.City})
	}
	if f.From != nil {
		where = append(where, sq.GtOrEq{"registration_date": *f.From})
	}
	if f.To != nil {
		where = append(where, sq.LtOrEq{"registration_date": *f.To})
	}

	return where
}

func buildPVZQuery(f model.PVZFilter) sq.SelectBuilder {
	qb := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("*").
		From("pvz").
		OrderBy("registration_date ASC")

	where := buildPVZWhere(f)
	if len(where) > 0 {
		qb = qb.Where(where)
	}

	if f.Limit != nil {
		qb = qb.Limit(uint64(*f.Limit))
	}

	if f.Page != nil && f.Limit != nil {
		qb = qb.Offset(uint64((*f.Page - 1) * (*f.Limit)))
	}

	return qb
}
