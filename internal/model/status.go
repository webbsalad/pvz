package model

import (
	"fmt"
	"strings"
)

type Status string

const (
	IN_PROGRESS Status = "in_progress"
	CLOSE       Status = "close"
)

func NewStatus(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "in_progress":
		return IN_PROGRESS, nil
	case "close":
		return CLOSE, nil
	default:
		return "", fmt.Errorf("unknown role: %s", s)
	}
}

func (s Status) String() string {
	switch s {
	case IN_PROGRESS:
		return "in_progress"
	case CLOSE:
		return "close"
	default:
		return ""
	}
}
