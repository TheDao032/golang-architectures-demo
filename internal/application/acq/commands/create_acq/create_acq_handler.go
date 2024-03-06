package createacq

import (
	"context"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/acq"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateACQHandler struct {
	logger logger.Logger
	cfg    *config.AppConfig
	repo   interfaces.ACQCommandRepository
}

func NewCreateACQHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.ACQCommandRepository,
) *CreateACQHandler {
	return &CreateACQHandler{logger, cfg, repo}
}

func (h *CreateACQHandler) Handle(ctx context.Context, createACQCommands []CreateACQCommand) ([]CreateACQCommand, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateACQHandler.Handle")
	defer span.Finish()

	for _, item := range createACQCommands {
		loc, _ := time.LoadLocation("UTC")
		acq := entities.ACQ{
			RxTime:       time.Unix(int64(item.Time), 0).In(loc),
			ExperimentId: item.ExperimentId,
			SignalId:     item.SignalId,
			Doppler:      item.Doppler,
			CodePhase:    item.CodePhase,
			AcfCorr:      item.AcfCorr,
			NoiseFloor:   item.NoiseFloor,
			Svid:         item.Svid,
			AcqMode:      item.AcqMode,
		}

		_, err := h.repo.CreateACQ(ctx, &acq)
		if err != nil {
			h.logger.Error(ctx, "ACQ has not been created successfully", zap.Error(err))
			return nil, err
		}
	}

	// loc, _ := time.LoadLocation("UTC")
	// acq := entities.ACQ{
	// 	RxTime:       time.Unix(int64(createACQCommand.RxTime), 0).In(loc),
	// 	ExperimentId: createACQCommand.ExperimentId,
	// 	SignalId:     createACQCommand.SignalId,
	// 	Doppler:      createACQCommand.Doppler,
	// 	CodePhase:    createACQCommand.CodePhase,
	// 	AcfCorr:      createACQCommand.AcfCorr,
	// 	NoiseFloor:   createACQCommand.NoiseFloor,
	// 	AcqMode:      createACQCommand.AcqMode,
	// }

	// _, err := h.repo.CreateACQ(ctx, &acq)

	// if err != nil {
	// 	h.logger.Error(ctx, "ACQ has not been created successfully", zap.Error(err))
	// 	return nil, err
	// }

	return createACQCommands, nil
}
