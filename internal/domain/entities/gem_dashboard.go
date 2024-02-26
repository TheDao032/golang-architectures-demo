// Include models that involved in business flow
package entities

import "database/sql"

type GemDashboard struct {
	Base
	UserId           string          `db:"user_id"`
	Pending          float64         `db:"pending"`
	Redeemable       float64         `db:"redeemable"`
	RedeemLimitation sql.NullFloat64 `db:"redeem_limitation"`
	Redeemed         float64         `db:"redeemed"`
	Status           sql.NullString  `db:"status"`
}
