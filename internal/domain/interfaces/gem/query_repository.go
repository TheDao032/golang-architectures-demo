package gem

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type GemQueryRepository interface {
	GetGemDashboard(ctx context.Context, userId string) (*entities.GemDashboard, error)
	GetGemSourcesByUserId(ctx context.Context, userId string, filter string) ([]entities.GemSource, error)
	GetGemSourceBySourceId(ctx context.Context, sourceId string) (*entities.GemSource, error)
	GetPendingGemSources(ctx context.Context) ([]entities.GemSource, error)
}
