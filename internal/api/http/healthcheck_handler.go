package http

import (
	"github.com/TheDao032/golang-architectures-demo/config"
	"github.com/TheDao032/golang-architectures-demo/database"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
)

type HealthcheckHandler struct {
	logger  logger.Logger
	cfg     *config.AppConfig
	readDb  *database.ReadDB
	writeDb *database.WriteDB
}

func NewHealthcheckHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	readDb *database.ReadDB,
	writeDb *database.WriteDB,
) *HealthcheckHandler {
	return &HealthcheckHandler{logger, cfg, readDb, writeDb}
}
