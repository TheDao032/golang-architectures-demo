package createraw

type CreateRAWCommand struct {
	Time         float64 `json:"rxTime"`
	ExperimentId int     `json:"experimentId"`
	SignalId     int     `json:"signalId"`
	Svid         int     `json:"svid"`
	FdRaw        float32 `json:"fdRaw"`
	FdRawRate    float32 `json:"fdRawRate"`
	PrRaw        float32 `json:"prRaw"`
}
