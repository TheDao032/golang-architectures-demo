package podservice

import (
	createpod "github.com/TheDao032/golang-architectures-demo/internal/application/pod/commands/create_pod"
)

type PODService struct {
	// Commands
	CreatePODHandler *createpod.CreatePODHandler

	// Queries
}

func NewPODService(
	// Commands
	createPODHandler *createpod.CreatePODHandler,
	// Queries
) *PODService {
	return &PODService{
		// Commands
		CreatePODHandler: createPODHandler,
		// Queries
	}
}
