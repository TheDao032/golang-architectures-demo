package createsta

import (
	"context"
	"database/sql"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/sta"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateSTAHandler struct {
	logger logger.Logger
	cfg    *config.AppConfig
	repo   interfaces.STACommandRepository
}

func NewCreateSTAHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.STACommandRepository,
) *CreateSTAHandler {
	return &CreateSTAHandler{logger, cfg, repo}
}

func (h *CreateSTAHandler) Handle(ctx context.Context, createSTACommands []CreateSTACommand) ([]CreateSTACommand, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateSTAHandler.Handle")
	defer span.Finish()

	for _, item := range createSTACommands {
		loc, _ := time.LoadLocation("UTC")
		var validApplication int16
		if item.OperationModeRunning != "" {
			validApplication = 3
		}

		sta := entities.STA{
			RxTime:             time.Unix(int64(item.Time), 0).In(loc),
			ExperimentId:       item.ExperimentId,
			RxOperationModeId:  sql.NullInt16{Int16: validApplication, Valid: item.OperationModeRunning != ""},
			SignalId:           item.SignalId,
			NumberOfChannels:   item.NumberOfChannels,
			Svid:               item.Svid,
			ChannelStatus:      item.ChannelStatus,
			NumberOfApps:       item.NumberOfApps,
			EccErrorCount:      item.EccErrorCount,
			CpuTemp:            item.CpuTemp,
			FrontendTemp:       item.FrontendTemp,
			Qn400VersionNumber: item.Qn400VersionNumber,
		}

		_, err := h.repo.CreateSTA(ctx, &sta)
		if err != nil {
			h.logger.Error(ctx, "STA has not been created successfully", zap.Error(err))
			return nil, err
		}
	}

	return createSTACommands, nil
}
