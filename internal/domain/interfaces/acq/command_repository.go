package acq

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type ACQCommandRepository interface {
	CreateACQ(ctx context.Context, acq *entities.ACQ) (bool, error)
}
