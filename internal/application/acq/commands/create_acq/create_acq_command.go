package createacq

type CreateACQCommand struct {
	RxTime       float64   `json:"rxTime"`
	ExperimentId int       `json:"experimentId"`
	SignalId     int       `json:"signalId"`
	Doppler      float32   `json:"doppler"`
	CodePhase    float32   `json:"codePhase"`
	AcfCorr      []float32 `json:"acfCorr`
	NoiseFloor   float32   `json:"noiseFloor"`
	AcqMode      int16     `json:"acqMode"`
}

type CreateACQResponse struct {
  Data interface{} `json:"data"`
}
