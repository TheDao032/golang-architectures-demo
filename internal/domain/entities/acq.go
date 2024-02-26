// Include models that involved in business flow
package entities

import "time"

type ACQ struct {
	RxTime       time.Time `db:"rx_time"`
	ExperimentId int       `db:"experiment_id"`
	SignalId     int       `db:"signal_id"`
	Doppler      float32   `db:"doppler"`
	CodePhase    float32   `db:"code_phase"`
	AcfCorr      []float32 `db:"acf_corr"`
	NoiseFloor   float32   `db:"noise_floor"`
	AcqMode      int16     `db:"acq_mode"`
}
