package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/webbsalad/pvz/internal/model"
)

func buildPVZWhere(pvz model.PVZ) (sq.And, error) {
	where := sq.And{}
	if pvz.City != "" {
		where = append(where, sq.Eq{"pvz_id": pvz.City})
	}

	return where, nil
}
