package getgemdashboard

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/dto"
	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
)

type GetGemDashboardHandler struct {
	logger logger.Logger
	repo   interfaces.GemQueryRepository
}

func NewGetGemDashboardHandler(
	logger logger.Logger,
	repo interfaces.GemQueryRepository,
) *GetGemDashboardHandler {
	return &GetGemDashboardHandler{logger, repo}
}

func (h *GetGemDashboardHandler) Handle(ctx context.Context, getGemQuery *GetGemDashboardQuery) (*GetGemDashboardResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetGemDashboardHandler.Handle")
	defer span.Finish()

	gem, err := h.repo.GetGemDashboard(ctx, getGemQuery.UserId)

	if err != nil {
		return nil, err
	}

	data := GetGemDashboard{
		Id:         gem.Id,
		UserId:     gem.UserId,
		Pending:    gem.Pending,
		Redeemable: gem.Redeemable,
		Redeemed:   gem.Redeemed,
	}

	response := &GetGemDashboardResponse{
		BaseResponse: dto.BaseResponse{
			Code:    0,
			Message: "",
			Success: true,
		},
		Data: data,
	}

	return response, nil
}
