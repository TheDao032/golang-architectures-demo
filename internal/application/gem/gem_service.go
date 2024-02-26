package gemservice

import (
	creategemdashboard "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_dashboard"
	updatependinggemdashboard "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/update_pending_gem_dashboard"

	creategemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_source"
	scanpendinggemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/scan_pending_gem_source"

	creategemhistory "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_history"

	getgemdashboard "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_gem_dashboard"
	getgemsourcebysource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_gem_source_by_source"
	getgemsourcebyuser "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_gem_source_by_user"
	getpendinggemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/queries/get_pending_gem_source"
)

type GemService struct {
	// Commands
	CreateGemDashboardHandler        *creategemdashboard.CreateGemDashboardHandler
	UpdatePendingGemDashboardHandler *updatependinggemdashboard.UpdatePendingGemDashboardHandler
	CreateGemSourceHandler           *creategemsource.CreateGemSourceHandler
	ScanPendingGemSourceHandler      *scanpendinggemsource.ScanPendingGemSourceHandler
	CreateGemHistoryHandler          *creategemhistory.CreateGemHistoryHandler

	// Queries
	GetGemDashboardHandler      *getgemdashboard.GetGemDashboardHandler
	GetGemSourceByUserHandler   *getgemsourcebyuser.GetGemSourceByUserHandler
	GetGemSourceBySourceHandler *getgemsourcebysource.GetGemSourceBySourceHandler
	GetPendingGemSourceHandler  *getpendinggemsource.GetPendingGemSourceHandler
}

func NewGemService(
	// Commands
	// createGemDashboardHandler *creategemdashboard.CreateGemDashboardHandler,
	updatePendingGemDashboardHandler *updatependinggemdashboard.UpdatePendingGemDashboardHandler,
	// createGemSourceHandler *creategemsource.CreateGemSourceHandler,
	scanPendingGemSourceHandler *scanpendinggemsource.ScanPendingGemSourceHandler,
	createGemHistoryHandler *creategemhistory.CreateGemHistoryHandler,

	// Queries
	getGemDashboardHandler *getgemdashboard.GetGemDashboardHandler,
	getGemSourceByUserHandler *getgemsourcebyuser.GetGemSourceByUserHandler,
	getGemSourceBySourceHandler *getgemsourcebysource.GetGemSourceBySourceHandler,
	getPendingGemSourceHandler *getpendinggemsource.GetPendingGemSourceHandler,
) *GemService {
	return &GemService{
		// Commands
		// CreateGemDashboardHandler:        createGemDashboardHandler,
		UpdatePendingGemDashboardHandler: updatePendingGemDashboardHandler,
		// CreateGemSourceHandler:           createGemSourceHandler,
		ScanPendingGemSourceHandler: scanPendingGemSourceHandler,
		CreateGemHistoryHandler:     createGemHistoryHandler,

		// Queries
		GetGemDashboardHandler:      getGemDashboardHandler,
		GetGemSourceByUserHandler:   getGemSourceByUserHandler,
		GetGemSourceBySourceHandler: getGemSourceBySourceHandler,
		GetPendingGemSourceHandler:  getPendingGemSourceHandler,
	}
}
