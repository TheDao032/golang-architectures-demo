package createpod

type CreatePODCommand struct {
	Time          float64   `json:"rxTime"`
	ExperimentId  int       `json:"experimentId"`
	ApplicationId int       `json:"applicationId"`
	Wn            int       `json:"wn"`
	Tow           int64     `json:"tow"`
	Decimals      float32   `json:"decimals"`
	NSat          int       `json:"nSat"`
	PosX          float32   `json:"posX"`
	PosY          float32   `json:"posY"`
	PosZ          float32   `json:"posZ"`
	VelX          float32   `json:"velX"`
	VelY          float32   `json:"velY"`
	VelZ          float32   `json:"velZ"`
	PosStd        []float32 `json:"posStd"`
	VelStd        []float32 `json:"velStd"`
	ClockBias     float32   `json:"clockBias"`
	ClockDrift    float32   `json:"clockDrift"`
	AmbigVec      []float32 `json:"ambigVec"`
	AmbigAcc      []float32 `json:"ambigAcc"`
}
