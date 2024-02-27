package createnav

import (
	"context"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/nav"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateNAVHandler struct {
	logger logger.Logger
	cfg    *config.AppConfig
	repo   interfaces.NAVCommandRepository
}

func NewCreateNAVHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.NAVCommandRepository,
) *CreateNAVHandler {
	return &CreateNAVHandler{logger, cfg, repo}
}

func (h *CreateNAVHandler) Handle(ctx context.Context, createNAVCommands []CreateNAVCommand) ([]CreateNAVCommand, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateNAVHandler.Handle")
	defer span.Finish()

	for _, item := range createNAVCommands {
		loc, _ := time.LoadLocation("UTC")
		nav := entities.NAV{
			Time:          time.Unix(int64(item.Time), 0).In(loc),
			ExperimentId:  item.ExperimentId,
			ApplicationId: 1,
			Wn:            item.Wn,
			Tow:           item.Tow,
			Decimals:      item.Decimals,
			NSat:          item.NSat,
			PosX:          item.PosX,
			PosY:          item.PosY,
			PosZ:          item.PosZ,
			VelX:          item.VelX,
			VelY:          item.VelY,
			VelZ:          item.VelZ,
			PosStd:        item.PosStd,
			VelStd:        item.VelStd,
			TimStd:        item.TimStd,
			ClockBias:     item.ClockBias,
			ClockDrift:    item.ClockDrift,
			Ggto:          item.Ggto,
			Gdop:          item.Gdop,
			Pdop:          item.Pdop,
			Hdop:          item.Hdop,
			Vdop:          item.Vdop,
			Tdop:          item.Tdop,
		}

		_, err := h.repo.CreateNAV(ctx, &nav)
		if err != nil {
			h.logger.Error(ctx, "NAV has not been created successfully", zap.Error(err))
			return nil, err
		}
	}

	return createNAVCommands, nil
}
