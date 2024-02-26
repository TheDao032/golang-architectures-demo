package getpendinggemsource

import (
	"time"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/dto"
)

type GetPendingGemSource struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	SourceId    string    `json:"sourceId"`
	Gems        float64   ` json:"gems"`
	Type        string    ` json:"type"`
	Status      string    ` json:"status"`
	Reason      string    ` json:"reason"`
	Metadata    string    ` json:"metadata"`
	CollectedAt time.Time ` json:"collectedAt"`
}

type GetPendingGemSourceResponse struct {
	dto.BaseResponse
	Data []*GetPendingGemSource `json:"data"`
}
