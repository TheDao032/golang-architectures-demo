package getgemdashboard

type GetGemDashboardQuery struct {
	UserId string `json:"userId" binding:"required"`
}
