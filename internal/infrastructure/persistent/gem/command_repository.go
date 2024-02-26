package gempersitent

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"
	"github.com/opentracing/opentracing-go"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type gemCommandRepository struct {
	writeDB database.WriteDB
	logger  logger.Logger
}

func NewGemCommandRepository(writeDb *database.WriteDB, logger logger.Logger) interfaces.GemCommandRepository {
	return &gemCommandRepository{*writeDb, logger}
}

func (repo *gemCommandRepository) CreateGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.CreateGemToDashboard")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO gem_dashboard(
			id,
			user_id,
			pending,
			redeemable,
			redeemed,
			status,
			created_by,
			created_at
		) VALUES (
			:id,
			:user_id,
			:pending,
			:redeemable,
			:redeemed,
			:status,
			:created_by,
			:created_at
		)`, gem,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *gemCommandRepository) UpdatePendingGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.UpdatePendingGemToDashboard")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		UPDATE gem_dashboard
		SET
			pending = :pending,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			user_id = :user_id`, gem,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *gemCommandRepository) UpdateRedeemableGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.UpdateRedeemableGemToDashboard")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		UPDATE gem_dashboard
		SET
			pending = :pending,
			redeemable = :redeemable,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			user_id = :user_id`, gem,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *gemCommandRepository) UpdateRedeemedGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.UpdateRedeemedGemToDashboard")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		UPDATE gem_dashboard
		SET
			redeemable = :redeemable,
			redeemed = :redeemed,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			user_id = :user_id`, gem,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *gemCommandRepository) CreateGemSource(ctx context.Context, gem *entities.GemSource) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.CreateGemSource")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO gem_source(
			id, 
			user_id, 
			source_id, 
			gems, 
			type, 
			status,
			reason,
			metadata, 
			collected_at, 
			created_at, 
			created_by 
		) VALUES (
			:id, 
			:user_id, 
			:source_id, 
			:gems, 
			:type, 
			:status,
			:reason,
			:metadata,
			:collected_at, 
			:created_at,
			:created_by
		)`,
		gem,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *gemCommandRepository) UpdateGemSource(ctx context.Context, gem *entities.GemSource) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.UpdateGemSource")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		UPDATE gem_source
		SET
			type = :type,
			status = :status,
			reason = :reason,
			metadata = :metadata,
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE
			source_id = :source_id`, gem,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem Source has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *gemCommandRepository) CreateGemHistory(ctx context.Context, gemHistory *entities.GemHistory) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemCommandRepository.CreateGemHistory")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO gem_history(
			id, 
			user_id, 
			source_id, 
			gems, 
			type, 
			status,
			reason,
			metadata, 
			collected_at,
			created_at, 
			created_by 
		) VALUES (
			:id, 
			:user_id, 
			:source_id, 
			:gems, 
			:type, 
			:status,
			:reason,
			:metadata,
			:collected_at,
			:created_at,
			:created_by
		)`, gemHistory,
	)

	if err != nil {
		repo.logger.Error(ctx, "Gem has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
