package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcPort = flag.Int("port", 3000, "gRPC port")
)

func NewServer(s *grpc.Server) grpc.ServiceRegistrar {
	return grpc.ServiceRegistrar(s)
}

func grpcOption() fx.Option {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", *grpcPort, err)
	}

	return fx.Options(
		fx.Provide(
			grpc.NewServer,
			NewServer,
		),
		fx.Invoke(
			func(s *grpc.Server, lc fx.Lifecycle) {
				reflection.Register(s)
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Println("starting grpc server")

						go func() {
							if err := s.Serve(lis); err != nil {
								log.Fatalf("failed to serve: %v", err)
							}
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						log.Println("stopping grpc server")
						s.GracefulStop()
						return nil
					},
				})
			},
		),
	)
}
