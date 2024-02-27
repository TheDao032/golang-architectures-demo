package nav

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type NAVCommandRepository interface {
	CreateNAV(ctx context.Context, nav *entities.NAV) (bool, error)
}
