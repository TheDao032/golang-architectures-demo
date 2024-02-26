package healthcheck

import (
	"context"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/database"
	"github.com/TheDao032/golang-architectures-demo/pkg/constants"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/heptiolabs/healthcheck"
)

func RunHealthCheck(
	logger logger.Logger,
	cfg *config.AppConfig,
	readDb *database.ReadDB,
	writeDb *database.WriteDB,
) healthcheck.Handler {
	ctx := context.Background()

	health := healthcheck.NewHandler()
	interval := time.Duration(cfg.Healthcheck.Interval) * time.Second

	livenessCheck(ctx, cfg, health)
	readinessCheck(ctx, logger, cfg, health, interval, readDb, writeDb)

	return health
}

func livenessCheck(ctx context.Context, cfg *config.AppConfig, health healthcheck.Handler) {
	health.AddLivenessCheck(constants.GoroutineThreshold, healthcheck.GoroutineCountCheck(cfg.Healthcheck.GoroutineThreshold))
}

func readinessCheck(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.AppConfig,
	health healthcheck.Handler,
	interval time.Duration,
	readDb *database.ReadDB,
	writeDb *database.WriteDB,
) {
	health.AddReadinessCheck(constants.ReadDatabase, healthcheck.AsyncWithContext(ctx, func() error {
		return (*readDb).Connection.PingContext(ctx)
	}, interval))

	health.AddReadinessCheck(constants.WriteDatabase, healthcheck.AsyncWithContext(ctx, func() error {
		return (*writeDb).Connection.PingContext(ctx)
	}, interval))

	// health.AddReadinessCheck(constants.Kafka, healthcheck.AsyncWithContext(ctx, func() error {
	// 	kafkaConn := kafka.ConnectKafkaBrokers(ctx, logger, &kafka.Config{
	// 		Brokers:  cfg.Kafka.Config.Brokers,
	// 		Username: cfg.Kafka.Config.Username,
	// 		Password: cfg.Kafka.Config.Password,
	// 	})
	// 	if _, err := kafkaConn.Brokers(); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }, interval))
}
