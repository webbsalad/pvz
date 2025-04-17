package model

import "time"

type Product struct {
	ID          ProductID
	ReceptionID ReceptionID
	Type        string

	DateTime time.Time
}

type ProductFilter struct {
	ReceptionID *ReceptionID
	Type        *string

	From *time.Time
	To   *time.Time
}
