package createnav

type CreateNAVCommand struct {
	Time          float64 `json:"rxTime"`
	ExperimentId  int     `json:"experimentId"`
	ApplicationId int     `json:"applicationId"`
	Wn            int     `json:"wn"`
	Tow           int64   `json:"tow"`
	Decimals      float32 `json:"decimals"`
	NSat          int     `json:"nSat"`
	PosX          float32 `json:"posX"`
	PosY          float32 `json:"posY"`
	PosZ          float32 `json:"posZ"`
	VelX          float32 `json:"velX"`
	VelY          float32 `json:"velY"`
	VelZ          float32 `json:"velZ"`
	PosStd        float32 `json:"posStd"`
	VelStd        float32 `json:"velStd"`
	TimStd        float32 `json:"timStd"`
	ClockBias     float32 `json:"clockBias"`
	ClockDrift    float32 `json:"clockDrift"`
	Ggto          float32 `json:"GGTO"`
	Gdop          float32 `json:"GDOP"`
	Pdop          float32 `json:"PDOP"`
	Hdop          float32 `json:"HDOP"`
	Vdop          float32 `json:"VDOP"`
	Tdop          float32 `json:"TDOP"`
}
