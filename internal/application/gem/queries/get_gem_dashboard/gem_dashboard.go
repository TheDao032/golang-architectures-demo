package getgemdashboard

import "github.com/TheDao032/golang-architectures-demo/internal/domain/dto"

type GetGemDashboard struct {
	Id               string  `json:"id"`
	UserId           string  `json:"userId"`
	Pending          float64 `json:"pending"`
	Redeemable       float64 `json:"redeemable"`
	RedeemLimitation float64 `json:"redeemLimitation"`
	Redeemed         float64 `json:"redeemed"`
	Status           string  `json:"status"`
}

type GetGemDashboardResponse struct {
	dto.BaseResponse
	Data GetGemDashboard `json:"data"`
}
