package database

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/config"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ReadDB struct {
	Connection *sqlx.DB
}

type WriteDB struct {
	Connection *sqlx.DB
}

func Open(cfg *config.DatabaseConfig, logger logger.Logger) (*ReadDB, *WriteDB) {
	ctx := context.Background()

	readDb, readErr := sqlx.Connect(cfg.ReadDbCfg.DbType, cfg.ReadDbCfg.ConnectionString)

	if readErr != nil {
		logger.Info(ctx, "ReadDb Connection String", zap.String("ConnectionString", cfg.ReadDbCfg.ConnectionString))
		logger.Info(ctx, "WriteDb Connection String", zap.String("ConnectionString", cfg.WriteDbCfg.ConnectionString))
		logger.Error(ctx, "Error Opening Read DB", zap.Error(readErr), zap.String("ReadDb Connection String", cfg.ReadDbCfg.ConnectionString), zap.String("WriteDb Connection String", cfg.WriteDbCfg.ConnectionString))
		panic(readErr)
	}
	readDB := &ReadDB{
		Connection: readDb,
	}

	writeDb, writeErr := sqlx.Connect(cfg.WriteDbCfg.DbType, cfg.WriteDbCfg.ConnectionString)
	if writeErr != nil {
		logger.Error(ctx, "Error Opening Write DB", zap.Error(writeErr))
		panic(writeErr)
	}

	writeDB := &WriteDB{
		Connection: writeDb,
	}

	logger.Info(ctx, "Connected to read & write database!!!")

	return readDB, writeDB
}
