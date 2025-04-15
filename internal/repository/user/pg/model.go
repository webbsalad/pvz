package pg

import "time"

type User struct {
	ID    string `db:"id"`
	Email string `db:"email"`
	Role  string `db:"role"`

	CreatedAt time.Time `db:"created_at"`
}
