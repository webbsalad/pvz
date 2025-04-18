package convertor

import (
	"errors"
	"log/slog"

	"github.com/webbsalad/pvz/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertError(err error, log *slog.Logger) error {
	switch {
	case errors.Is(err, model.ErrUnauthenticated):
		log.Warn("unauthenticated", slog.String("err", err.Error()))
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, model.ErrPermissionDenied):
		log.Warn("permission denied", slog.String("err", err.Error()))
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, model.ErrNotFound):
		log.Warn("not found", slog.String("err", err.Error()))
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, model.ErrAlreadyExist):
		log.Warn("already exist", slog.String("err", err.Error()))
		return status.Error(codes.AlreadyExists, err.Error())
	}

	log.Error("internal server error", slog.String("err", err.Error()))
	return status.Errorf(codes.Internal, "internal server error: %v", err)
}
