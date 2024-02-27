package stapersitent

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/sta"
	"github.com/opentracing/opentracing-go"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type staCommandRepository struct {
	writeDB database.WriteDB
	logger  logger.Logger
}

func NewSTACommandRepository(writeDb *database.WriteDB, logger logger.Logger) interfaces.STACommandRepository {
	return &staCommandRepository{*writeDb, logger}
}

func (repo *staCommandRepository) CreateSTA(ctx context.Context, sta *entities.STA) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "staCommandRepository.CreateSTA")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO sta(
      time,
			experiment_id,
      rx_operation_mode_id,
			signal_id,
      number_of_channels,
			sv_id,
      channel_status,
			number_of_apps,
			ecc_error_count,
			cpu_temp,
			frontend_temp,
			qn400_version_number
		) VALUES (
      :time,
      :experiment_id,
      :rx_operation_mode_id,
      :signal_id,
      :number_of_channels,
      :sv_id,
      :channel_status,
      :number_of_apps,
      :ecc_error_count,
      :cpu_temp,
      :frontend_temp,
      :qn400_version_number
		)`, sta,
	)

	if err != nil {
		repo.logger.Error(ctx, "STA has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
