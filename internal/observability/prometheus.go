package observability

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitMetrics(metricsAddr string) func() error {
	srv := &http.Server{Addr: metricsAddr}
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv.Handler = mux
	go func() {
		_ = srv.ListenAndServe()
	}()

	return func() error {
		return srv.Shutdown(context.Background())
	}
}
