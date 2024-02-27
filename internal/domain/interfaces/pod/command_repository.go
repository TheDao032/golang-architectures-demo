package pod

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type PODCommandRepository interface {
	CreatePOD(ctx context.Context, nav *entities.POD) (bool, error)
}
