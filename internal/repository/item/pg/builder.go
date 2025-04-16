package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/webbsalad/pvz/internal/model"
)

func buildReceptionWhere(reception model.Reception) (sq.And, error) {
	where := sq.And{}
	if reception.PVZID.String() != "" {
		where = append(where, sq.Eq{"pvz_id": reception.PVZID.String()})
	}

	if reception.Status.String() != "" {
		where = append(where, sq.Eq{"status": reception.Status.String()})
	}

	return where, nil
}
