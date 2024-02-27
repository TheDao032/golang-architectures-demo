package createraw

import (
	"context"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/raw"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateRAWHandler struct {
	logger logger.Logger
	cfg    *config.AppConfig
	repo   interfaces.RAWCommandRepository
}

func NewCreateRAWHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.RAWCommandRepository,
) *CreateRAWHandler {
	return &CreateRAWHandler{logger, cfg, repo}
}

func (h *CreateRAWHandler) Handle(ctx context.Context, createRAWCommands []CreateRAWCommand) ([]CreateRAWCommand, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateRAWHandler.Handle")
	defer span.Finish()

	for _, item := range createRAWCommands {
		loc, _ := time.LoadLocation("UTC")
		raw := entities.RAW{
			RxTime:       time.Unix(int64(item.Time), 0).In(loc),
			ExperimentId: item.ExperimentId,
			SignalId:     item.SignalId,
			Svid:         item.Svid,
			FdRaw:        item.FdRaw,
			FdRawRate:    item.FdRawRate,
			PrRaw:        item.PrRaw,
		}

		_, err := h.repo.CreateRAW(ctx, &raw)
		if err != nil {
			h.logger.Error(ctx, "RAW has not been created successfully", zap.Error(err))
			return nil, err
		}
	}

	return createRAWCommands, nil
}
