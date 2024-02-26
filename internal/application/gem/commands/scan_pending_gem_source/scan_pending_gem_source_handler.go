package scanpendinggemsource

import (
	"context"
	"database/sql"
	"strconv"
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

type ScanPendingGemSourceHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	queryRepo     interfaces.GemQueryRepository
	commandRepo   interfaces.GemCommandRepository
	kafkaProducer *k.Producer
}

func NewScanPendingGemSourceHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	queryRepo interfaces.GemQueryRepository,
	commandRepo interfaces.GemCommandRepository,
	kafkaProducer *k.Producer,
) *ScanPendingGemSourceHandler {
	return &ScanPendingGemSourceHandler{logger, cfg, queryRepo, commandRepo, kafkaProducer}
}

func (h *ScanPendingGemSourceHandler) Handle(ctx context.Context) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ScanPendingGemSourceHandler.Handle")
	defer span.Finish()

	listPendingGemSource, err := h.queryRepo.GetPendingGemSources(ctx)

	if err != nil {
		return false, err
	}

	for _, pendingGemSource := range listPendingGemSource {
		// getGemSource, err := h.queryRepo.GetGemSourceBySourceId(ctx, pendingGemSource.SourceId)
		getGemDashboard, err := h.queryRepo.GetGemDashboard(ctx, pendingGemSource.UserId)

		if getGemDashboard.Base.Id == "" {
			return false, err
		}

		gemSourceBaseEntity := entities.Base{
			Id:        pendingGemSource.Id,
			UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedBy: sql.NullString{String: pendingGemSource.UpdatedBy.String, Valid: true},
		}

		pendingGemSource.Reason.Valid = true
		pendingGemSource.Metadata.Valid = true
		pendingGemSource.CollectedAt.Valid = true

		gemSource := entities.GemSource{
			Base:        gemSourceBaseEntity,
			UserId:      pendingGemSource.UserId,
			SourceId:    pendingGemSource.SourceId,
			Gems:        pendingGemSource.Gems,
			Type:        pendingGemSource.Type,
			Reason:      pendingGemSource.Reason,
			Metadata:    pendingGemSource.Metadata,
			Status:      entities.TransStatusDeposited,
			CollectedAt: pendingGemSource.CollectedAt,
		}

		updateGemSourceSuccess, err := h.commandRepo.UpdateGemSource(ctx, &gemSource)

		if err != nil || !updateGemSourceSuccess {
			h.logger.Error(ctx, "Update Gem Source Failed", zap.Error(err))
			return false, err
		}

		createGemHistoryBaseEntity := entities.Base{
			Id:        uuid.New().String(),
			CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			CreatedBy: sql.NullString{String: gemSource.Base.CreatedBy.String, Valid: true},
		}
		gemHistory := &entities.GemHistory{
			Base:        createGemHistoryBaseEntity,
			SourceId:    gemSource.SourceId,
			UserId:      gemSource.UserId,
			Gems:        gemSource.Gems,
			Type:        gemSource.Type,
			Status:      entities.TransStatusDeposited,
			Reason:      gemSource.Reason,
			Metadata:    gemSource.Metadata,
			CollectedAt: gemSource.CollectedAt,
		}

		_, err = h.commandRepo.CreateGemHistory(ctx, gemHistory)

		if err != nil {
			h.logger.Error(ctx, "Gem history has not been created successfully", zap.Error(err))
			return false, err
		}

		gemDashboardBaseEntity := entities.Base{
			Id:        getGemDashboard.Id,
			UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedBy: sql.NullString{String: gemSource.Base.CreatedBy.String, Valid: true},
		}

		float64PendingGems, err := strconv.ParseFloat(strconv.FormatFloat(getGemDashboard.Pending-gemSource.Gems, 'f', 2, 64), 64)
		float64RedeemableGems, err := strconv.ParseFloat(strconv.FormatFloat(getGemDashboard.Redeemable+gemSource.Gems, 'f', 2, 64), 64)

		if err != nil {
			h.logger.Error(ctx, "Error while converting pending gems to float64", zap.Error(err))
			return false, err
		}

		gemDashboard := entities.GemDashboard{
			Base:       gemDashboardBaseEntity,
			UserId:     pendingGemSource.UserId,
			Pending:    float64PendingGems,
			Redeemable: float64RedeemableGems,
		}
		_, err = h.commandRepo.UpdateRedeemableGemToDashboard(ctx, &gemDashboard)

		if err != nil {
			h.logger.Error(ctx, "Update Redeemable Failed", zap.Error(err))
			return false, err
		}
	}

	if err != nil {
		h.logger.Error(ctx, "Source has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
