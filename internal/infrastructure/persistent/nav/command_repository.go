package navpersitent

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/nav"
	"github.com/opentracing/opentracing-go"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type navCommandRepository struct {
	writeDB database.WriteDB
	logger  logger.Logger
}

func NewNAVCommandRepository(writeDb *database.WriteDB, logger logger.Logger) interfaces.NAVCommandRepository {
	return &navCommandRepository{*writeDb, logger}
}

func (repo *navCommandRepository) CreateNAV(ctx context.Context, nav *entities.NAV) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "navCommandRepository.CreateNAV")
	defer span.Finish()

	_, err := repo.writeDB.Connection.NamedExec(`
		INSERT INTO nav(
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
			tim_std,
			clock_bias,
			clock_drift,
			ggto,
			gdop,
			pdop,
			hdop,
			vdop,
			tdop
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
			:tim_std,
			:clock_bias,
			:clock_drift,
			:ggto,
			:gdop,
			:pdop,
			:hdop,
			:vdop,
			:tdop
		)`, nav,
	)

	if err != nil {
		repo.logger.Error(ctx, "NAV has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
