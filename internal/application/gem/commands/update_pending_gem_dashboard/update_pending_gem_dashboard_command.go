package updatependinggemdashboard

type UpdatePendingGemDashboardCommand struct {
	Id               string  `json:"id" binding:"required"`
	UserId           string  `json:"userId" binding:"required"`
	Pending          float64 `json:"pending"`
	Redeemable       float64 `json:"redeemable"`
	RedeemLimitation float64 `json:"redeemLimitation"`
	Redeemed         float64 `json:"redeemed"`
	Status           string  `json:"status" binding:"required"`
}
