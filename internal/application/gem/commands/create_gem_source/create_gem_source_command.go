package creategemsource

import "time"

type CreateGemSourceCommand struct {
	Id          string    `json:"id" binding:"required"`
	UserId      string    `json:"userId" binding:"required"`
	SourceId    string    `json:"sourceId" binding:"required"`
	Gems        float64   `json:"gems" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	CollectedAt time.Time `json:"collectedAt" binding:"required"`
	Reason      string    `json:"reason"`
	Metadata    string    `json:"metadata"`
	CreatedBy   string    `json:"createdBy"`
}
