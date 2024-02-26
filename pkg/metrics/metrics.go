package metrics

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/config"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func RunMetrics(logger logger.Logger, cfg *config.AppConfig) func() {
	return func() {
		gin.SetMode(gin.ReleaseMode)
		metricsServer := gin.New()

		metricsServer.GET(cfg.Metrics.PrometheusPath, prometheusHandler())

		ctx := context.Background()
		logger.Info(ctx, "Metrics server is running on port", zap.String("Metrics port", cfg.Metrics.PrometheusPort))
		if err := metricsServer.Run(cfg.Metrics.PrometheusPort); err != nil {
			// If service uses both of http & grpc, it probaly happens error here(already bind the same port)
			// It's still good to go
			logger.Error(ctx, "metricsServer.Run", zap.Error(err))
		}
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
