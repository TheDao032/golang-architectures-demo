package raw

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type RAWCommandRepository interface {
	CreateRAW(ctx context.Context, acq *entities.RAW) (bool, error)
}
