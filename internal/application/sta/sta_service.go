package staservice

import (
	createsta "github.com/TheDao032/golang-architectures-demo/internal/application/sta/commands/create_sta"
)

type STAService struct {
	// Commands
	CreateSTAHandler *createsta.CreateSTAHandler

	// Queries
}

func NewSTAService(
	// Commands
	createSTAHandler *createsta.CreateSTAHandler,
	// Queries
) *STAService {
	return &STAService{
		// Commands
		CreateSTAHandler: createSTAHandler,
		// Queries
	}
}
