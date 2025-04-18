package metadata

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const HTTPCodeHeader = "http-code"

func SetHTTPStatus(ctx context.Context, code int) error {
	md := metadata.Pairs(HTTPCodeHeader, fmt.Sprintf("%d", code))
	return grpc.SendHeader(ctx, md)
}
