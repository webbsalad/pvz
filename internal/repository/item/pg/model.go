package pg

import "time"

type Reception struct {
	ID     string `db:"id"`
	PVZID  string `db:"pvz_id"`
	Status string `db:"status"`

	DateTime time.Time `db:"date_time"`
}

type Product struct {
	ID          string `db:"id"`
	ReceptionID string `db:"reception_id"`
	Type        string `db:"type"`

	DateTime time.Time `db:"date_time"`
}
