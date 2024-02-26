package gem

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/entities"
)

type GemCommandRepository interface {
	UpdatePendingGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error)
	UpdateRedeemableGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error)
	UpdateRedeemedGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error)
	CreateGemToDashboard(ctx context.Context, gem *entities.GemDashboard) (bool, error)

	CreateGemSource(ctx context.Context, gem *entities.GemSource) (bool, error)
	UpdateGemSource(ctx context.Context, gem *entities.GemSource) (bool, error)

	CreateGemHistory(ctx context.Context, gem *entities.GemHistory) (bool, error)
}
