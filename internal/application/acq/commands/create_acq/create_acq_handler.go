package createacq

import (
	"context"
	"database/sql"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/acq"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateACQdHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	repo          interfaces.ACQCommandRepository
}

func NewCreateGemDashboardHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.ACQCommandRepository,
) *CreateACQHandler {
	return &CreateACQdHandler{logger, cfg, repo}
}

func (h *CreateACQdHandler) Handle(ctx context.Context, createACQCommand *CreateACQCommand) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateACQHandler.Handle")
	defer span.Finish()

	acq := entities.ACQ{
    RxTime: time.Duration(createACQCommand.RxTime) * float64(time.Second),
    ExperimentId: createACQCommand.ExperimentId,
	}

	_, err := h.repo.CreateACQ(ctx, &gem)

	if err != nil {
		h.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
