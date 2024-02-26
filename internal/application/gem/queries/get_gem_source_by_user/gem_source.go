package getgemsourcebyuser

import (
	"time"

	"github.com/TheDao032/golang-architectures-demo/internal/domain/dto"
)

type GetGemSourceByUser struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	SourceId    string    `json:"sourceId"`
	Gems        float64   ` json:"gems"`
	Type        string    ` json:"type"`
	Status      string    ` json:"status"`
	Reason      string    ` json:"reason"`
	Metadata    string    ` json:"metadata"`
	CollectedAt time.Time ` json:"collectedAt"`
	DefinedAt   time.Time ` json:"definedAt"`
	DefinedBy   string    ` json:"definedBy"`
	ModifiedAt  time.Time ` json:"modifiedAt"`
	ModifiedBy  string    ` json:"modifiedBy"`
}

type GetGemSourceByUserResponse struct {
	dto.BaseResponse
	Data []GetGemSourceByUser `json:"data"`
}
