//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/TheDao032/golang-architectures-demo/config"
	database "github.com/TheDao032/golang-architectures-demo/database"
	"github.com/TheDao032/golang-architectures-demo/internal/api"
	"github.com/TheDao032/golang-architectures-demo/internal/api/http"

	// creategemdashboard "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_dashboard"
	// creategemhistory "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_history"
	// creategemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_source"
	// scanpendinggemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/scan_pending_gem_source"
	// updatependinggemdashboard "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/update_pending_gem_dashboard"
	// getgemdashboard "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_gem_dashboard"
	// getgemsourcebysource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_gem_source_by_source"
	// getgemsourcebyuser "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_gem_source_by_user"
	// getpendinggemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_pending_gem_source"

	// gemrepo "github.com/TheDao032/golang-architectures-demo/internal/infrastructure/persistent/gem"
	// service "github.com/TheDao032/golang-architectures-demo/internal/service"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/google/wire"
)

var container = wire.NewSet(
	api.NewApiContainer,
)

var apiSet = wire.NewSet(
	http.NewServer,
	// scheduler.NewScheduler,
)

var serviceSet = wire.NewSet(
	// service.NewService,
	http.NewHealthcheckHandler,
)

// var specificServiceSet = wire.NewSet(
// 	gemservice.NewGemService,
// )

// var handlerSet = wire.NewSet(
// creategemdashboard.NewCreateGemDashboardHandler,
// updatependinggemdashboard.NewUpdatePendingGemDashboardHandler,
// creategemsource.NewCreateGemSourceHandler,
// scanpendinggemsource.NewScanPendingGemSourceHandler,
// creategemhistory.NewCreateGemHistoryHandler,

// getgemdashboard.NewGetGemDashboardHandler,
// getgemsourcebysource.NewGetGemSourceBySourceHandler,
// getgemsourcebyuser.NewGetGemSourceByUserHandler,
// getpendinggemsource.NewGetPendingGemSourceHandler,
// )

// var repoSet = wire.NewSet(
// 	gemrepo.NewGemQueryRepository,
// 	gemrepo.NewGemCommandRepository,
// )

func InitializeContainer(
	appCfg *config.AppConfig,
	logger logger.Logger,
	readDB *database.ReadDB,
	writeDB *database.WriteDB,
) *api.ApiContainer {
	wire.Build(serviceSet, apiSet, container)
	return &api.ApiContainer{}
}
