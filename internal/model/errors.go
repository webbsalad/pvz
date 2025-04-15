package model

import (
	"errors"
	"fmt"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExist     = errors.New("already exist")
	ErrUnauthenticated  = errors.New("unauthenticated")
)

var (
	ErrUserAlreadyExist = fmt.Errorf("user already exist: %w", ErrAlreadyExist)
	ErrPasswordMismatch = fmt.Errorf("password mismatch: %w", ErrUnauthenticated)
	ErrPVZAlreadyExist  = fmt.Errorf("pvz already exist: %w", ErrAlreadyExist)
)
