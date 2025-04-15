package model

import (
	"fmt"
	"strings"
)

type Role string

const (
	CLIENT    Role = "client"
	MODERATOR Role = "moderator"
	EMPLOYEE  Role = "employee"
)

func NewRole(s string) (Role, error) {
	switch strings.ToLower(s) {
	case "client":
		return CLIENT, nil
	case "moderator":
		return MODERATOR, nil
	case "employee":
		return EMPLOYEE, nil
	default:
		return "", fmt.Errorf("unknown role: %s", s)
	}
}

func (r Role) String() string {
	switch r {
	case CLIENT:
		return "client"
	case MODERATOR:
		return "moderator"
	case EMPLOYEE:
		return "employee"
	default:
		return ""
	}
}
