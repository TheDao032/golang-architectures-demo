package creategemsource

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

type CreateGemSourceHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	queryRepo     interfaces.GemQueryRepository
	commandRepo   interfaces.GemCommandRepository
	kafkaProducer *k.Producer
}

func NewCreateGemSourceHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	queryRepo interfaces.GemQueryRepository,
	commandRepo interfaces.GemCommandRepository,
	kafkaProducer *k.Producer,
) *CreateGemSourceHandler {
	return &CreateGemSourceHandler{logger, cfg, queryRepo, commandRepo, kafkaProducer}
}

func (h *CreateGemSourceHandler) Handle(ctx context.Context, createGemSourceCommand *CreateGemSourceCommand) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateGemSourceHandler.Handle")
	defer span.Finish()

	createGemSourceBaseEntity := entities.Base{
		Id:        uuid.New().String(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		CreatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
	}

	gemSourceMessage := entities.GemSource{
		Base:        createGemSourceBaseEntity,
		UserId:      createGemSourceCommand.UserId,
		SourceId:    createGemSourceCommand.SourceId,
		Gems:        createGemSourceCommand.Gems,
		Type:        createGemSourceCommand.Type,
		Reason:      sql.NullString{String: createGemSourceCommand.Reason, Valid: true},
		Metadata:    sql.NullString{String: createGemSourceCommand.Metadata, Valid: true},
		CollectedAt: sql.NullTime{Time: createGemSourceCommand.CollectedAt, Valid: true},
		Status:      entities.TransStatusPending,
	}

	if gemSourceMessage.Type == entities.TransTypeReverse {

		getGemSource, err := h.queryRepo.GetGemSourceBySourceId(ctx, gemSourceMessage.SourceId)

		if err != nil {
			return false, err
		}

		getGemDashboard, err := h.queryRepo.GetGemDashboard(ctx, getGemSource.UserId)
		if getGemDashboard.Base.Id == "" {
			h.logger.Error(ctx, "Dashboard info is null", zap.Error(err))
			return false, err
		}

		if getGemSource.Type == gemSourceMessage.Type {
			h.logger.Error(ctx, "Revert transaction already has one", zap.Error(err))
			return false, nil
		}

		if getGemSource.Status == entities.TransStatusPending {

			updateGemSourceBaseEntity := entities.Base{
				Id:        getGemSource.Id,
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}

			gemSourcePendingReverted := &entities.GemSource{
				Base:     updateGemSourceBaseEntity,
				UserId:   getGemSource.UserId,
				SourceId: gemSourceMessage.SourceId,
				Status:   entities.TransStatusReverted,
				Type:     gemSourceMessage.Type,
				Reason:   getGemSource.Reason,
				Metadata: getGemSource.Metadata,
			}

			_, err := h.commandRepo.UpdateGemSource(ctx, gemSourcePendingReverted)

			if err != nil {
				h.logger.Error(ctx, "Transaction has not been updated successfully", zap.Error(err))
				return false, err
			}

			createGemHistoryBaseEntity := entities.Base{
				Id:        uuid.New().String(),
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				CreatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}
			gemHistoryPendingReverted := &entities.GemHistory{
				Base:        createGemHistoryBaseEntity,
				UserId:      getGemSource.UserId,
				SourceId:    gemSourceMessage.SourceId,
				Gems:        getGemSource.Gems,
				Type:        gemSourceMessage.Type,
				Status:      entities.TransStatusReverted,
				Reason:      getGemSource.Reason,
				Metadata:    getGemSource.Metadata,
				CollectedAt: sql.NullTime{Time: time.Now(), Valid: true},
			}

			createGemHistorySuccess, err := h.commandRepo.CreateGemHistory(ctx, gemHistoryPendingReverted)

			if err != nil || !createGemHistorySuccess {
				h.logger.Error(ctx, "Gem history has not been created successfully", zap.Error(err))
				return false, err
			}

			updateDashboardBaseEntity := entities.Base{
				Id:        getGemDashboard.Id,
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}

			float64PendingGems, err := strconv.ParseFloat(strconv.FormatFloat(getGemDashboard.Pending-getGemSource.Gems, 'f', 2, 64), 64)
			if err != nil {
				h.logger.Error(ctx, "Error while converting pending gems to float64", zap.Error(err))
				return false, err
			}

			gemDashboard := entities.GemDashboard{
				Base:    updateDashboardBaseEntity,
				UserId:  getGemSource.UserId,
				Pending: float64PendingGems,
			}
			updateGemDashboard, err := h.commandRepo.UpdatePendingGemToDashboard(ctx, &gemDashboard)

			if err != nil || !updateGemDashboard {
				h.logger.Error(ctx, "Dashboard has not been updated successfully", zap.Error(err))
				return false, err
			}
		} else {
			updateGemSourceBaseEntity := entities.Base{
				Id:        getGemSource.Id,
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}

			gemSourceRedeemableReverted := entities.GemSource{
				Base:     updateGemSourceBaseEntity,
				UserId:   gemSourceMessage.UserId,
				SourceId: gemSourceMessage.SourceId,
				Gems:     getGemSource.Gems,
				Type:     gemSourceMessage.Type,
				Reason:   getGemSource.Reason,
				Metadata: getGemSource.Metadata,
				Status:   entities.TransStatusReverted,
			}

			updateGemSourceSuccess, err := h.commandRepo.UpdateGemSource(ctx, &gemSourceRedeemableReverted)

			if err != nil || !updateGemSourceSuccess {
				h.logger.Error(ctx, "Transaction has not been updated successfully", zap.Error(err))
				return false, err
			}

			createGemHistoryBaseEntity := entities.Base{
				Id:        uuid.New().String(),
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				CreatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}
			gemHistoryRedeemableReverted := &entities.GemHistory{
				Base:        createGemHistoryBaseEntity,
				UserId:      getGemSource.UserId,
				SourceId:    gemSourceMessage.SourceId,
				Gems:        getGemSource.Gems,
				Type:        gemSourceMessage.Type,
				Status:      entities.TransStatusReverted,
				Reason:      getGemSource.Reason,
				Metadata:    getGemSource.Metadata,
				CollectedAt: sql.NullTime{Time: time.Now(), Valid: true},
			}

			createGemHistorySuccess, err := h.commandRepo.CreateGemHistory(ctx, gemHistoryRedeemableReverted)

			if err != nil || !createGemHistorySuccess {
				h.logger.Error(ctx, "Transaction has not been created successfully", zap.Error(err))
				return false, err
			}

			updateGemDashboardBaseEntity := entities.Base{
				Id:        getGemDashboard.Id,
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}

			float64RedeemableGems, err := strconv.ParseFloat(strconv.FormatFloat(getGemDashboard.Redeemable-getGemSource.Gems, 'f', 2, 64), 64)
			if err != nil {
				h.logger.Error(ctx, "Error while converting pending gems to float64", zap.Error(err))
				return false, err
			}

			gemDashboard := entities.GemDashboard{
				Base:       updateGemDashboardBaseEntity,
				UserId:     getGemSource.UserId,
				Pending:    getGemDashboard.Pending,
				Redeemable: float64RedeemableGems,
			}
			updateGemDashboardSuccess, err := h.commandRepo.UpdateRedeemableGemToDashboard(ctx, &gemDashboard)

			if err != nil || !updateGemDashboardSuccess {
				h.logger.Error(ctx, "Transaction has not been updated successfully", zap.Error(err))
				return false, err
			}

		}
	} else {
		getGemDashboard, err := h.queryRepo.GetGemDashboard(ctx, createGemSourceCommand.UserId)

		if err != nil {
			h.logger.Error(ctx, "Getting error while get dashboard information", zap.Error(err))
			return false, err
		}

		if getGemDashboard.Base.Id == "" {
			createGemSourceSuccess, err := h.commandRepo.CreateGemSource(ctx, &gemSourceMessage)

			if err != nil || !createGemSourceSuccess {
				h.logger.Error(ctx, "Transaction has not been created successfully", zap.Error(err))
				return false, err
			}

			createGemHistoryBaseEntity := entities.Base{
				Id:        uuid.New().String(),
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				CreatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}
			gemHistoryPending := &entities.GemHistory{
				Base:        createGemHistoryBaseEntity,
				UserId:      gemSourceMessage.UserId,
				SourceId:    gemSourceMessage.SourceId,
				Gems:        gemSourceMessage.Gems,
				Type:        gemSourceMessage.Type,
				Status:      entities.TransStatusPending,
				Reason:      gemSourceMessage.Reason,
				Metadata:    gemSourceMessage.Metadata,
				CollectedAt: gemSourceMessage.CollectedAt,
			}

			createGemHistorySuccess, err := h.commandRepo.CreateGemHistory(ctx, gemHistoryPending)

			if err != nil || !createGemHistorySuccess {
				h.logger.Error(ctx, "Transaction history has not been created successfully", zap.Error(err))
				return false, err
			}

			createorUpdateDashboardBaseEntity := entities.Base{
				Id:        uuid.New().String(),
				CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				CreatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
				UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
				UpdatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
			}

			float64PendingGems, err := strconv.ParseFloat(strconv.FormatFloat(gemSourceMessage.Gems, 'f', 2, 64), 64)
			if err != nil {
				h.logger.Error(ctx, "Error while converting pending gems to float64", zap.Error(err))
				return false, err
			}

			gemDashboard := entities.GemDashboard{
				Base:     createorUpdateDashboardBaseEntity,
				UserId:   createGemSourceCommand.UserId,
				Redeemed: 0,
				Pending:  float64PendingGems,
			}

			createGemDashboard, err := h.commandRepo.CreateGemToDashboard(ctx, &gemDashboard)
			if err != nil || !createGemDashboard {
				h.logger.Error(ctx, "Gem dashboard has not been created successfully", zap.Error(err))
				return false, err
			}

			getGemDashboard, err = h.queryRepo.GetGemDashboard(ctx, createGemSourceCommand.UserId)

			if err != nil {
				h.logger.Error(ctx, "Getting error while get dashboard information", zap.Error(err))
				return false, err
			}

			return true, nil
		}

		createGemSourceSuccess, err := h.commandRepo.CreateGemSource(ctx, &gemSourceMessage)

		if err != nil || !createGemSourceSuccess {
			h.logger.Error(ctx, "Transaction has not been created successfully", zap.Error(err))
			return false, err
		}

		createGemHistoryBaseEntity := entities.Base{
			Id:        uuid.New().String(),
			CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			CreatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
		}
		gemHistoryPending := &entities.GemHistory{
			Base:        createGemHistoryBaseEntity,
			UserId:      gemSourceMessage.UserId,
			SourceId:    gemSourceMessage.SourceId,
			Gems:        gemSourceMessage.Gems,
			Type:        gemSourceMessage.Type,
			Status:      entities.TransStatusPending,
			Reason:      gemSourceMessage.Reason,
			Metadata:    gemSourceMessage.Metadata,
			CollectedAt: gemSourceMessage.CollectedAt,
		}

		createGemHistorySuccess, err := h.commandRepo.CreateGemHistory(ctx, gemHistoryPending)

		if err != nil || !createGemHistorySuccess {
			h.logger.Error(ctx, "Transaction history has not been created successfully", zap.Error(err))
			return false, err
		}

		updateDashboardBaseEntity := entities.Base{
			Id:        getGemDashboard.Id,
			UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedBy: sql.NullString{String: createGemSourceCommand.CreatedBy, Valid: true},
		}

		float64PendingGems, err := strconv.ParseFloat(strconv.FormatFloat(getGemDashboard.Pending+gemSourceMessage.Gems, 'f', 2, 64), 64)
		if err != nil {
			h.logger.Error(ctx, "Error while converting pending gems to float64", zap.Error(err))
			return false, err
		}

		gemDashboard := entities.GemDashboard{
			Base:    updateDashboardBaseEntity,
			UserId:  gemSourceMessage.UserId,
			Pending: float64PendingGems,
		}

		updateGemDashboardSuccess, err := h.commandRepo.UpdatePendingGemToDashboard(ctx, &gemDashboard)

		if err != nil || !updateGemDashboardSuccess {
			h.logger.Error(ctx, "Transaction has not been updated successfully", zap.Error(err))
			return false, err
		}
	}

	return true, nil
}
