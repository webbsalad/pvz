package model

import (
	"errors"
	"fmt"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExist     = errors.New("already exist")
	ErrInvalidArgument  = errors.New("invalid argument")
)

var (
	ErrUserAlreadyExist      = fmt.Errorf("user already exist: %w", ErrAlreadyExist)
	ErrPasswordMismatch      = fmt.Errorf("password mismatch: %w", ErrInvalidArgument)
	ErrPVZAlreadyExist       = fmt.Errorf("pvz already exist: %w", ErrAlreadyExist)
	ErrJwtExpired            = fmt.Errorf("expired jwt token: %w", ErrPermissionDenied)
	ErrWrongRole             = fmt.Errorf("wrong role: %w", ErrPermissionDenied)
	ErrReceptionAlreadyExist = fmt.Errorf("reception already exist : %w", ErrAlreadyExist)
	ErrReceptionNotFound     = fmt.Errorf("reception not found: %w", ErrNotFound)
	ErrProductNotFound       = fmt.Errorf("product not found: %w", ErrNotFound)
	ErrPVZNotFound           = fmt.Errorf("pvz not found; %w", ErrNotFound)
	ErrUserNotFound          = fmt.Errorf("user not found: %w", ErrNotFound)
)
