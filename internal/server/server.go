package server

import (
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yuxsr/notificator/internal/server/service"
	protov1 "github.com/yuxsr/yuxsr-dev-pb/gencode/go/proto/v1"
	"google.golang.org/grpc"
)

func RegisterNewGRPCServer(config service.NotificatorServiceConfig) *grpc.Server {
	// Setup metrics.
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)
	// exemplarFromContext := func(ctx context.Context) prometheus.Labels {
	// 	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
	// 		return prometheus.Labels{"traceID": span.TraceID().String()}
	// 	}
	// 	return nil
	// }

	notificatorService := service.NewNotificatorService(config)
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
			// recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)
	protov1.RegisterNotificatorServiceServer(s, notificatorService)
	srvMetrics.InitializeMetrics(s)
	return s
}
