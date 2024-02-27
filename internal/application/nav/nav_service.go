package navservice

import (
	createnav "github.com/TheDao032/golang-architectures-demo/internal/application/nav/commands/create_nav"
)

type NAVService struct {
	// Commands
	CreateNAVHandler *createnav.CreateNAVHandler

	// Queries
}

func NewNAVService(
	// Commands
	createNAVHandler *createnav.CreateNAVHandler,
	// Queries
) *NAVService {
	return &NAVService{
		// Commands
		CreateNAVHandler: createNAVHandler,
		// Queries
	}
}
