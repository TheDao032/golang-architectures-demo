package createacq

type CreateACQCommand struct {
	RxTime       float64   `json:"rxTime" binding:"required"`
	ExperimentId int       `json:"experimentId" binding:"required"`
	SignalId     int       `json:"signalId" binding:"required"`
	Doppler      float32   `json:"doppler"`
	CodePhase    float32   `json:"codePhase"`
	AcfCorr      []float32 `json:"acfCorr`
	NoiseFloor   float32   `json:"noiseFloor" binding:"required"`
	AcqMode      int16     `json:"acqMode" binding:"required"`
}

type CreateACQResponse struct {
  Data interface{} `json:"data"`
}
