package metadata

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func GetBarerToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata")
	}

	token := md.Get("bearer_token")
	if len(token) == 0 {
		return "", fmt.Errorf("missing token")
	}

	return token[0], nil
}
