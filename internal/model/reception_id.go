package model

import (
	"fmt"

	"github.com/google/uuid"
)

type ReceptionID uuid.UUID

func (id ReceptionID) String() string {
	return uuid.UUID(id).String()
}

func NewReceptionID(s string) (ReceptionID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ReceptionID{}, fmt.Errorf("parse: %w", err)
	}
	return ReceptionID(id), nil
}
