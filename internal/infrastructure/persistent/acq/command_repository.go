package acqpersitent

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/acq"
	"github.com/opentracing/opentracing-go"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type acqCommandRepository struct {
	writeDB database.WriteDB
	logger  logger.Logger
}

func NewACQCommandRepository(writeDb *database.WriteDB, logger logger.Logger) interfaces.ACQCommandRepository {
	return &acqCommandRepository{*writeDb, logger}
}

func (repo *acqCommandRepository) CreateACQ(ctx context.Context, acq *entities.ACQ) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "acqCommandRepository.CreateACQ")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO acq(
      time,
			experiment_id,
			signal_id,
			doppler,
			code_phase,
			acf_corr,
			noise_floor,
			acq_mode
		) VALUES (
			:time,
			:experiment_id,
			:signal_id,
			:doppler,
			:code_phase,
			:acf_corr,
			:noise_floor,
      :acq_mode
		)`, acq,
	)

	if err != nil {
		repo.logger.Error(ctx, "ACQ has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
