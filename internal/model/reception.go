package model

import "time"

type Reception struct {
	ID     ReceptionID
	PVZID  PVZID
	Status Status

	DateTime time.Time
}
