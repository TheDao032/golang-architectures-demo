package startup

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/database"
	"github.com/TheDao032/golang-architectures-demo/internal"
	"github.com/TheDao032/golang-architectures-demo/internal/api"
	"github.com/TheDao032/golang-architectures-demo/pkg/metrics"

	"github.com/TheDao032/go-backend-utils-architecture/command"
	"github.com/TheDao032/go-backend-utils-architecture/localization"
	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/TheDao032/go-backend-utils-architecture/validation"
	"github.com/gammazero/workerpool"
)

func runServer(
	ctx context.Context,
	logger logger.Logger,
	container *api.ApiContainer,
	readDb *database.ReadDB,
	writeDb *database.WriteDB,
	// kafkaConn *k.Conn,
) {
	// Run scheduler
	// container.Scheduler.Run()

	wp := workerpool.New(2)
	// Run metrics
	wp.Submit(metrics.RunMetrics(logger, cfg))
	// Run Http server
	wp.Submit(container.HttpServer.Run)

	wp.StopWait()
}

func registerDependencies(ctx context.Context, logger logger.Logger) (*api.ApiContainer, *database.ReadDB, *database.WriteDB) {
	// Open database connection
	readDb, writeDb := database.Open(cfg.Database, logger)
	// Register kafka
	// kafkaProducer := kafka.NewProducer(ctx, (*kafka.Config)(cfg.Kafka.Config), logger, &k.Writer{})
	// Register dependencies
	return internal.InitializeContainer(cfg, logger, readDb, writeDb),
		readDb,
		writeDb
}

var cfg *config.AppConfig

func Execute() {
	// Init AppConfig
	cfg = &config.AppConfig{}

	// Init commands
	command.UseCommands(
		command.WithStartCommand(start, cfg, "DATABASE.WRITEDB"),
		// command.WithMigrationCommand("DATABASE.WRITEDB"),
	)
}

func start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	// Init logger
	logger := logger.NewZapLogger(cfg.Logger.LogLevel)
	// Register dependencies
	container, readDb, writeDb := registerDependencies(ctx, logger)
	// Init resources for localization
	localization.InitResources(cfg.Http.Resources)
	// Init kakfa
	// producerConfigs := []kafka.ProducerConfig{
	// 	{
	// 		TopicName:         cfg.Kafka.Producers.GemCreateProducer.TopicName,
	// 		InitTopic:         cfg.Kafka.Producers.GemCreateProducer.InitTopic,
	// 		NumPartitions:     cfg.Kafka.Producers.GemCreateProducer.NumPartitions,
	// 		ReplicationFactor: cfg.Kafka.Producers.GemCreateProducer.ReplicationFactor,
	// 	},
	// }
	// consumerConfigs := []kafka.ConsumerConfig{
	// 	{
	// 		GroupId:   cfg.Kafka.Consumers.GemCreateConsumer.GroupId,
	// 		TopicName: cfg.Kafka.Consumers.GemCreateConsumer.TopicName,
	// 		NumWorker: cfg.Kafka.Consumers.GemCreateConsumer.NumWorker,
	// 		Worker:    container.Consumer.ProcessMessages,
	// 	},
	// }
	// kafkaConn := kafka.UseKafka(
	// 	ctx,
	// 	logger,
	// 	&kafka.Config{
	// 		Brokers:  cfg.Kafka.Config.Brokers,
	// 		Username: cfg.Kafka.Config.Username,
	// 		Password: cfg.Kafka.Config.Password,
	// 	},
	// 	producerConfigs,
	// 	consumerConfigs,
	// )
	// Init validation
	validation.UseValidation()
	// Run server
	runServer(ctx, logger, container, readDb, writeDb)
}
