package podpersitent

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/pod"
	"github.com/opentracing/opentracing-go"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type podCommandRepository struct {
	writeDB database.WriteDB
	logger  logger.Logger
}

func NewPODCommandRepository(writeDb *database.WriteDB, logger logger.Logger) interfaces.PODCommandRepository {
	return &podCommandRepository{*writeDb, logger}
}

func (repo *podCommandRepository) CreatePOD(ctx context.Context, pod *entities.POD) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "podCommandRepository.CreatePOD")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO pod(
      time,
			experiment_id,
			application_id,
			wn,
			tow,
			decimals,
			n_sat,
			pos_x,
			pos_y,
			pos_z,
			vel_x,
			vel_y,
			vel_z,
			pos_std,
			vel_std,
			clock_bias,
			clock_drift,
			ambig_vec,
			ambig_acc
		) VALUES (
			:time,
			:experiment_id,
			:application_id,
      :wn,
			:tow,
			:decimals,
			:n_sat,
			:pos_x,
			:pos_y,
			:pos_z,
			:vel_x,
			:vel_y,
			:vel_z,
			:pos_std,
			:vel_std,
			:clock_bias,
			:clock_drift,
			:ambig_vec,
			:ambig_acc
		)`, pod,
	)

	if err != nil {
		repo.logger.Error(ctx, "POD has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
