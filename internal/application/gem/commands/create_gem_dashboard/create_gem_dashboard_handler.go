package creategemdashboard

import (
	"context"
	"database/sql"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	k "github.com/TheDao032/go-backend-utils-architecture/kafka"
	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateGemDashboardHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	repo          interfaces.GemCommandRepository
	kafkaProducer *k.Producer
}

func NewCreateGemDashboardHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.GemCommandRepository,
	kafkaProducer *k.Producer,
) *CreateGemDashboardHandler {
	return &CreateGemDashboardHandler{logger, cfg, repo, kafkaProducer}
}

func (h *CreateGemDashboardHandler) Handle(ctx context.Context, createGemCommand *CreateGemDashboardCommand) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateGemDashboardHandler.Handle")
	defer span.Finish()

	base := entities.Base{
		Id: uuid.New().String(),
	}

	gem := entities.GemDashboard{
		Base:             base,
		UserId:           createGemCommand.UserId,
		Pending:          createGemCommand.Pending,
		Redeemable:       createGemCommand.Redeemable,
		Redeemed:         createGemCommand.Redeemed,
		RedeemLimitation: sql.NullFloat64{Float64: createGemCommand.RedeemLimitation, Valid: createGemCommand.RedeemLimitation != 0},
	}

	_, err := h.repo.CreateGemToDashboard(ctx, &gem)

	if err != nil {
		h.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
