// Include models that involved in business flow
package entities

import (
	"database/sql"
	"time"
)

type STA struct {
	RxTime             time.Time     `db:"time"`
	ExperimentId       int           `db:"experiment_id"`
	RxOperationModeId  sql.NullInt16 `db:"rx_operation_mode_id"`
	SignalId           int           `db:"signal_id"`
	NumberOfChannels   int16         `db:"number_of_channels"`
	Svid               int           `db:"sv_id"`
	ChannelStatus      int16         `db:"channel_status"`
	NumberOfApps       int16         `db:"number_of_apps"`
	EccErrorCount      int64         `db:"ecc_error_count"`
	CpuTemp            float32       `db:"cpu_temp"`
	FrontendTemp       float32       `db:"frontend_temp"`
	Qn400VersionNumber int64         `db:"qn400_version_number"`
}
