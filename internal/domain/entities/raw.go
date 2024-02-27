// Include models that involved in business flow
package entities

import (
	"time"
)

type RAW struct {
	RxTime       time.Time `db:"time"`
	ExperimentId int       `db:"experiment_id"`
	SignalId     int       `db:"signal_id"`
	Svid         int       `db:"sv_id"`
	FdRaw        float32   `db:"fd_raw"`
	FdRawRate    float32   `db:"fd_rate_raw"`
	PrRaw        float32   `db:"pr_raw"`
}
