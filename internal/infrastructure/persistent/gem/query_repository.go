package gempersitent

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/TheDao032/golang-architectures-demo/database"
	entities "github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
)

type gemQueryRepository struct {
	readDB database.ReadDB
}

func NewGemQueryRepository(readDb *database.ReadDB) interfaces.GemQueryRepository {
	return &gemQueryRepository{*readDb}
}

func (repo *gemQueryRepository) GetGemDashboard(ctx context.Context, userId string) (*entities.GemDashboard, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemQueryRepository.GetGemDashboard")
	defer span.Finish()

	gem := entities.GemDashboard{}
	err := repo.readDB.Connection.GetContext(ctx, &gem, `
		SELECT 
			id,
			user_id,
			pending,
			redeemable,
			redeem_limitation,
			redeemed,
			created_at,
			created_by,
			updated_by, 
			updated_at 
		FROM gem_dashboard WHERE user_id=$1`, userId)

	if err == sql.ErrNoRows {
		return &gem, nil
	}

	return &gem, err
}

func (repo *gemQueryRepository) GetGemSourcesByUserId(ctx context.Context, userId string, filter string) ([]entities.GemSource, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemQueryRepository.GetGemSourcesByUserId")
	defer span.Finish()

	gems := []entities.GemSource{}

	var condition string = ""
	if filter != "" {
		condition = "user_id=$1 AND status=$2"
	} else {
		condition = "user_id=$1 AND status!=$2"
	}

	rows, err := repo.readDB.Connection.Queryx(fmt.Sprintf(`
		SELECT 
			id,
			user_id,
			source_id,
			gems,
			type,
			status,
			metadata,
			reason,
			collected_at,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM gem_source 
		WHERE %v
		ORDER BY updated_at DESC`, condition), userId, filter)

	for rows.Next() {
		gem := entities.GemSource{}

		err := rows.StructScan(&gem)
		if err != nil {
			// handle error
			return gems, err
		}
		gems = append(gems, gem)
	}

	return gems, err
}

func (repo *gemQueryRepository) GetGemSourceBySourceId(ctx context.Context, sourceId string) (*entities.GemSource, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemQueryRepository.GetGemSourceBySourceId")
	defer span.Finish()

	gem := entities.GemSource{}
	err := repo.readDB.Connection.GetContext(ctx, &gem, `
		SELECT 
			id,
			user_id,
			source_id,
			gems,
			type,
			status,
			metadata,
			reason,
			collected_at,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM gem_source WHERE source_id=$1`, sourceId)

	return &gem, err
}

func (repo *gemQueryRepository) GetPendingGemSources(ctx context.Context) ([]entities.GemSource, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "gemQueryRepository.GetPendingGemSources")
	defer span.Finish()

	gems := []entities.GemSource{}
	rows, err := repo.readDB.Connection.Queryx(`
		SELECT 
			id,
			user_id,
			source_id,
			gems,
			type,
			status,
			reason,
			metadata,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM 
			gem_source
		WHERE 
			status='pending' AND
			to_char(collected_at, 'DD/MM/YYYY') = to_char(now(), 'DD/MM/YYYY')`)

	for rows.Next() {
		gem := entities.GemSource{}

		err := rows.StructScan(&gem)
		if err != nil {
			// handle error
			return gems, err
		}
		gems = append(gems, gem)
	}

	return gems, err
}
