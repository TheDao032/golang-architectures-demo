package sta

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type STACommandRepository interface {
	CreateSTA(ctx context.Context, acq *entities.STA) (bool, error)
}
