package creategemhistory

import (
	"context"
	"database/sql"
	"time"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	k "github.com/TheDao032/go-backend-utils-architecture/kafka"
	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateGemHistoryHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	repo          interfaces.GemCommandRepository
	kafkaProducer *k.Producer
}

func NewCreateGemHistoryHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.GemCommandRepository,
	kafkaProducer *k.Producer,
) *CreateGemHistoryHandler {
	return &CreateGemHistoryHandler{logger, cfg, repo, kafkaProducer}
}

func (h *CreateGemHistoryHandler) Handle(ctx context.Context, createGemHistoryCommand *CreateGemHistoryCommand) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdateGem.Handle")
	defer span.Finish()

	base := entities.Base{
		Id:        uuid.New().String(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		CreatedBy: sql.NullString{String: createGemHistoryCommand.CreatedBy, Valid: true},
	}

	gem := entities.GemHistory{
		Base:   base,
		UserId: createGemHistoryCommand.UserId,
		Gems:   createGemHistoryCommand.Gems,
		Type:   createGemHistoryCommand.Type,
		Reason: sql.NullString{String: createGemHistoryCommand.Reason, Valid: true},
		Status: createGemHistoryCommand.Status,
	}

	_, err := h.repo.CreateGemHistory(ctx, &gem)

	if err != nil {
		h.logger.Error(ctx, "Gem has not been updated successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
