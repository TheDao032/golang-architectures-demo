package getpendinggemsource

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/dto"

	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
)

type GetPendingGemSourceHandler struct {
	logger logger.Logger
	repo   interfaces.GemQueryRepository
}

func NewGetPendingGemSourceHandler(
	logger logger.Logger,
	repo interfaces.GemQueryRepository,
) *GetPendingGemSourceHandler {
	return &GetPendingGemSourceHandler{logger, repo}
}

func (h *GetPendingGemSourceHandler) Handle(ctx context.Context, getGemQuery *GetPendingGemSourceQuery) (*GetPendingGemSourceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetPendingGemSourceHandler.Handle")
	defer span.Finish()

	gems := []*GetPendingGemSource{}
	gem, err := h.repo.GetGemSourcesByUserId(ctx, getGemQuery.UserId, "")

	if err != nil {
		return nil, err
	}

	for _, g := range gem {
		p := GetPendingGemSource{
			Id:          g.Id,
			UserId:      g.UserId,
			SourceId:    g.SourceId,
			Gems:        g.Gems,
			Type:        g.Type,
			Status:      g.Status,
			Reason:      g.Reason.String,
			Metadata:    g.Metadata.String,
			CollectedAt: g.CollectedAt.Time,
		}
		gems = append(gems, &p)
	}

	response := &GetPendingGemSourceResponse{
		BaseResponse: dto.BaseResponse{
			Code:    0,
			Message: "",
			Success: true,
		},
		Data: gems,
	}

	return response, nil
}
