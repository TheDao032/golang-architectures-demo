package createsta

type CreateSTACommand struct {
	Time                 float64 `json:"rxTime"`
	ExperimentId         int     `json:"experimentId"`
	OperationModeRunning string  `json:"operationModeRunning"`
	RxOperationModeId    int     `json:"rxOperationModeId"`
	SignalId             int     `json:"signalId"`
	NumberOfChannels     int16   `json:"numberOfChannels"`
	Svid                 int     `json:"svid"`
	ChannelStatus        int16   `json:"channelStatus"`
	NumberOfApps         int16   `json:"numberOfApps"`
	EccErrorCount        int64   `json:"eccErrorCount"`
	CpuTemp              float32 `json:"cpuTemp"`
	FrontendTemp         float32 `json:"frontendTemp"`
	Qn400VersionNumber   int64   `json:"qn400VersionNumber"`
}
