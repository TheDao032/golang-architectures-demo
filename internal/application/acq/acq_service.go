package acqservice

import (
	createacq "github.com/TheDao032/golang-architectures-demo/internal/application/acq/commands/create_acq"
)

type ACQService struct {
	// Commands
	CreateACQHandler        *createacq.CreateACQdHandler

	// Queries
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
