package updatependinggemdashboard

import (
	"context"
	"database/sql"

	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	k "github.com/TheDao032/go-backend-utils-architecture/kafka"
	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type UpdatePendingGemDashboardHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	repo          interfaces.GemCommandRepository
	kafkaProducer *k.Producer
}

func NewUpdatePendingGemDashboardHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.GemCommandRepository,
	kafkaProducer *k.Producer,
) *UpdatePendingGemDashboardHandler {
	return &UpdatePendingGemDashboardHandler{logger, cfg, repo, kafkaProducer}
}

func (h *UpdatePendingGemDashboardHandler) Handle(ctx context.Context, updateGemCommand *UpdatePendingGemDashboardCommand) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdatePendingGemDashboardHandler.Handle")
	defer span.Finish()

	base := entities.Base{
		Id: updateGemCommand.Id,
	}

	gem := entities.GemDashboard{
		Base:             base,
		UserId:           updateGemCommand.UserId,
		Pending:          updateGemCommand.Pending,
		Redeemable:       updateGemCommand.Redeemable,
		Redeemed:         updateGemCommand.Redeemed,
		RedeemLimitation: sql.NullFloat64{Float64: updateGemCommand.RedeemLimitation, Valid: updateGemCommand.RedeemLimitation != 0},
	}

	_, err := h.repo.UpdatePendingGemToDashboard(ctx, &gem)

	if err != nil {
		h.logger.Error(ctx, "Gem has not been updated successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
