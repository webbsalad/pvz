package pg

import "time"

type PVZ struct {
	ID   string `db:"id"`
	City string `db:"city"`

	RegistrationDate time.Time `db:"registration_date"`
}
