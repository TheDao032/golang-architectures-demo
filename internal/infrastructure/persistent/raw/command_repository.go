package rawpersitent

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/raw"
	"github.com/opentracing/opentracing-go"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type rawCommandRepository struct {
	writeDB database.WriteDB
	logger  logger.Logger
}

func NewRAWCommandRepository(writeDb *database.WriteDB, logger logger.Logger) interfaces.RAWCommandRepository {
	return &rawCommandRepository{*writeDb, logger}
}

func (repo *rawCommandRepository) CreateRAW(ctx context.Context, raw *entities.RAW) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "rawCommandRepository.CreateRAW")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO raw(
      time,
			experiment_id,
			signal_id,
			sv_id,
			fd_raw,
			fd_rate_raw,
			pr_raw
		) VALUES (
			:time,
			:experiment_id,
			:signal_id,
      :sv_id,
      :fd_raw,
      :fd_rate_raw,
      :pr_raw
		)`, raw,
	)

	if err != nil {
		repo.logger.Error(ctx, "ACQ has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
