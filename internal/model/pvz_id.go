package model

import (
	"fmt"

	"github.com/google/uuid"
)

type PVZID uuid.UUID

func (id PVZID) String() string {
	return uuid.UUID(id).String()
}

func NewPVZID(s string) (PVZID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return PVZID{}, fmt.Errorf("parse: %w", err)
	}
	return PVZID(id), nil
}
