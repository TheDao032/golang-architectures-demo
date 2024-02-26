package getgemsourcebysource

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/dto"

	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
)

type GetGemSourceBySourceHandler struct {
	logger logger.Logger
	repo   interfaces.GemQueryRepository
}

func NewGetGemSourceBySourceHandler(
	logger logger.Logger,
	repo interfaces.GemQueryRepository,
) *GetGemSourceBySourceHandler {
	return &GetGemSourceBySourceHandler{logger, repo}
}

func (h *GetGemSourceBySourceHandler) Handle(ctx context.Context, getGemQuery *GetGemSourceBySourceQuery) (*GetGemSourceBySourceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetGemSourceBySourceHandler.Handle")
	defer span.Finish()

	// Get from cache first
	gem, err := h.repo.GetGemSourceBySourceId(ctx, getGemQuery.SourceId)

	if err != nil {
		return nil, err
	}

	data := GetGemSourceBySource{
		Id:          gem.Id,
		UserId:      gem.UserId,
		SourceId:    gem.SourceId,
		Gems:        gem.Gems,
		Type:        gem.Type,
		Status:      gem.Status,
		Reason:      gem.Reason.String,
		Metadata:    gem.Metadata.String,
		CollectedAt: gem.CollectedAt.Time,
	}

	response := &GetGemSourceBySourceResponse{
		BaseResponse: dto.BaseResponse{
			Code:    0,
			Message: "",
			Success: true,
		},
		Data: data,
	}

	return response, nil
}
