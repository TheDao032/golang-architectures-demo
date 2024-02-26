package entities

import "database/sql"

const (
	TransTypeEarning    = "deposit"
	TransTypeRedemption = "redeem"
	TransTypeReverse    = "revert"
)

const (
	TransStatusPending   = "pending"
	TransStatusDeposited = "deposited"
	TransStatusRedeemed  = "redeemed"
	TransStatusReverted  = "reverted"
)

type Base struct {
	Id        string         `db:"id"`
	CreatedAt sql.NullTime   `db:"created_at"`
	CreatedBy sql.NullString `db:"created_by"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	UpdatedBy sql.NullString `db:"updated_by"`
}
