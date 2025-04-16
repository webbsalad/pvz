package model

import (
	"fmt"

	"github.com/google/uuid"
)

type ProductID uuid.UUID

func (id ProductID) String() string {
	return uuid.UUID(id).String()
}

func NewProductID(s string) (ProductID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ProductID{}, fmt.Errorf("parse: %w", err)
	}
	return ProductID(id), nil
}
