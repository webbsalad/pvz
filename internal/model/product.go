package model

import "time"

type Product struct {
	ID          ProductID
	ReceptionID ReceptionID
	Type        string

	DateTime time.Time
}
