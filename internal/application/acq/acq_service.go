package acqservice

import (
	createacq "github.com/TheDao032/golang-architectures-demo/internal/application/acq/commands/create_acq"
)

type ACQService struct {
	// Commands
	CreateACQHandler *createacq.CreateACQHandler

	// Queries
}

func NewACQService(
	// Commands
	createACQHandler *createacq.CreateACQHandler,
	// Queries
) *ACQService {
	return &ACQService{
		// Commands
		CreateACQHandler: createACQHandler,
		// Queries
	}
}
