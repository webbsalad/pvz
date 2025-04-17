package model

import (
	"time"
)

type PVZ struct {
	ID   PVZID
	City string

	RegistrationDate time.Time
}

type PVZFilter struct {
	IDs  []PVZID
	City *string

	Page  *int32
	Limit *int32

	From *time.Time
	To   *time.Time
}
