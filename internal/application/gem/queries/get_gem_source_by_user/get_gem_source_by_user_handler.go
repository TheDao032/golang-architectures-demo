package getgemsourcebyuser

import (
	"context"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/dto"

	interfaces "github.com/TheDao032/golang-architectures-demo/internal/domain/interfaces/gem"

	"github.com/TheDao032/go-backend-utils-architecture/logger"
	"github.com/opentracing/opentracing-go"
)

type GetGemSourceByUserHandler struct {
	logger logger.Logger
	repo   interfaces.GemQueryRepository
}

func NewGetGemSourceByUserHandler(
	logger logger.Logger,
	repo interfaces.GemQueryRepository,
) *GetGemSourceByUserHandler {
	return &GetGemSourceByUserHandler{logger, repo}
}

func (h *GetGemSourceByUserHandler) Handle(ctx context.Context, getGemQuery *GetGemSourceByUserQuery) (*GetGemSourceByUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetGemSourceByUserHandler.Handle")
	defer span.Finish()

	// Get from cache first
	gems := []GetGemSourceByUser{}

	gem, err := h.repo.GetGemSourcesByUserId(ctx, getGemQuery.UserId, getGemQuery.Filter)

	if err != nil {
		return nil, err
	}

	for _, g := range gem {
		p := GetGemSourceByUser{
			Id:          g.Id,
			UserId:      g.UserId,
			SourceId:    g.SourceId,
			Gems:        g.Gems,
			Type:        g.Type,
			Status:      g.Status,
			Reason:      g.Reason.String,
			Metadata:    g.Metadata.String,
			CollectedAt: g.CollectedAt.Time,
			DefinedAt:   g.CreatedAt.Time,
			DefinedBy:   g.CreatedBy.String,
			ModifiedAt:  g.UpdatedAt.Time,
			ModifiedBy:  g.UpdatedBy.String,
		}
		gems = append(gems, p)
	}

	response := &GetGemSourceByUserResponse{
		BaseResponse: dto.BaseResponse{
			Code:    0,
			Message: "",
			Success: true,
		},
		Data: gems,
	}

	return response, nil
}
