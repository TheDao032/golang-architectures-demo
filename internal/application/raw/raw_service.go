package rawservice

import (
	createraw "github.com/TheDao032/golang-architectures-demo/internal/application/raw/commands/create_raw"
)

type RAWService struct {
	// Commands
	CreateRAWHandler *createraw.CreateRAWHandler

	// Queries
}

func NewRAWService(
	// Commands
	createRAWHandler *createraw.CreateRAWHandler,
	// Queries
) *RAWService {
	return &RAWService{
		// Commands
		CreateRAWHandler: createRAWHandler,
		// Queries
	}
}
