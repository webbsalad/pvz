package model

import (
	"fmt"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

func NewUserID(s string) (UserID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UserID{}, fmt.Errorf("parse: %w", err)
	}
	return UserID(id), nil
}
