package model

import (
	"time"
)

type User struct {
	ID    UserID
	Email string
	Role  Role

	CreatedAt time.Time
}
