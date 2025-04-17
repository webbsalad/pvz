package model

import (
	"time"
)

type Reception struct {
	ID     ReceptionID
	PVZID  PVZID
	Status Status

	DateTime time.Time
}

type ReceptionFilter struct {
	PVZID  *PVZID
	Status *Status
	From   *time.Time
	To     *time.Time
}
