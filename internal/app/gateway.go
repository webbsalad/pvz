package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	pb "github.com/webbsalad/pvz/internal/pb/github.com/webbsalad/pvz/pvz_v1"
)

var (
	gwPort = flag.Int("gw_port", 8080, "HTTP port")
)

func newCORSHandler(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(h)
}

func newBearerAuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			r.Header.Set("Grpc-Metadata-bearer_token", token)
		}
		h.ServeHTTP(w, r)
	})
}

func gatewayOption() fx.Option {
	flag.Parse()

	return fx.Invoke(func(lc fx.Lifecycle) {
		mux := runtime.NewServeMux()

		opts := []grpc.DialOption{grpc.WithInsecure()}

		err := pb.RegisterPVZServiceHandlerFromEndpoint(
			context.Background(),
			mux,
			fmt.Sprintf("localhost:%d", *grpcPort),
			opts,
		)
		if err != nil {
			log.Fatalf("failed register gateway: %v", err)
		}

		handler := newBearerAuthHandler(newCORSHandler(mux))

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", *gwPort),
			Handler: handler,
		}

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("failed http server starting: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("stopping HTTP gateway")
				return srv.Shutdown(ctx)
			},
		})
	})
}
