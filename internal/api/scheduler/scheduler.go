package scheduler

import (
	"context"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/service"
	"go.uber.org/zap"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	logger  logger.Logger
	cfg     *config.AppConfig
	service *service.Service
}

func NewScheduler(
	logger logger.Logger,
	cfg *config.AppConfig,
	s *service.Service,
) *Scheduler {
	return &Scheduler{logger, cfg, s}
}

func (s *Scheduler) Run() {
	go s.start()
}

func (sch *Scheduler) start() {
	s := gocron.NewScheduler(time.UTC)

	s.Cron(sch.cfg.Scheduler.CronExpression).Do(sch.scanPendingSourceList)
	s.StartAsync()
}

func (sch *Scheduler) scanPendingSourceList() {
	ctx := context.Background()
	sch.logger.Info(ctx, "Ready To Scan Pending Source", zap.String("Scheduler Config", sch.cfg.Scheduler.CronExpression))
	// sch.service.GemService.ScanPendingGemSourceHandler.Handle(ctx)
}
