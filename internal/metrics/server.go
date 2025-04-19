package metrics

import (
	"context"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle) {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{Addr: ":9000", Handler: mux}

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("metrics HTTP server error: %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown(ctx)
			},
		})
	})
}
