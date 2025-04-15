package model

import "time"

type PVZ struct {
	ID   PVZID
	City string

	RegistrationDate time.Time
}
