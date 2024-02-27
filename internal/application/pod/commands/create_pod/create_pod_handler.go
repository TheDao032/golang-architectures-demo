package createpod

import (
	"context"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/pod"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreatePODHandler struct {
	logger logger.Logger
	cfg    *config.AppConfig
	repo   interfaces.PODCommandRepository
}

func NewCreatePODHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.PODCommandRepository,
) *CreatePODHandler {
	return &CreatePODHandler{logger, cfg, repo}
}

func (h *CreatePODHandler) Handle(ctx context.Context, createPODCommands []CreatePODCommand) ([]CreatePODCommand, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreatePODHandler.Handle")
	defer span.Finish()

	for _, item := range createPODCommands {
		loc, _ := time.LoadLocation("UTC")
		pod := entities.POD{
			Time:          time.Unix(int64(item.Time), 0).In(loc),
			ExperimentId:  item.ExperimentId,
			ApplicationId: 2,
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
			ClockBias:     item.ClockBias,
			ClockDrift:    item.ClockDrift,
			AmbigVec:      item.AmbigVec,
			AmbigAcc:      item.AmbigAcc,
		}

		_, err := h.repo.CreatePOD(ctx, &pod)
		if err != nil {
			h.logger.Error(ctx, "POD has not been created successfully", zap.Error(err))
			return nil, err
		}
	}

	return createPODCommands, nil
}
