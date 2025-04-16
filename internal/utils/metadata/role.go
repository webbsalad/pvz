package metadata

import (
	"context"
	"fmt"

	"github.com/webbsalad/pvz/internal/model"
	"google.golang.org/grpc/metadata"
)

func GetRole(ctx context.Context) (model.Role, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata")
	}

	metaRole := md.Get("role")
	if len(metaRole) == 0 {
		return "", fmt.Errorf("missing role")
	}

	role, err := model.NewRole(metaRole[0])
	if err != nil {
		return "", fmt.Errorf("convert meta role to model: %w", err)
	}

	return role, nil
}
