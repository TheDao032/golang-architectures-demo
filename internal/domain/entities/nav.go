// Include models that involved in business flow
package entities

import (
	"time"
)

type NAV struct {
	Time          time.Time `db:"time"`
	ExperimentId  int       `db:"experiment_id"`
	ApplicationId int       `db:"application_id"`
	Wn            int       `db:"wn"`
	Tow           int64     `db:"tow"`
	Decimals      float32   `db:"decimals"`
	NSat          int       `db:"n_sat"`
	PosX          float32   `db:"pos_x"`
	PosY          float32   `db:"pos_y"`
	PosZ          float32   `db:"pos_z"`
	VelX          float32   `db:"vel_x"`
	VelY          float32   `db:"vel_y"`
	VelZ          float32   `db:"vel_z"`
	PosStd        float32   `db:"pos_std"`
	VelStd        float32   `db:"vel_std"`
	TimStd        float32   `db:"tim_std"`
	ClockBias     float32   `db:"clock_bias"`
	ClockDrift    float32   `db:"clock_drift"`
	Ggto          float32   `db:"ggto"`
	Gdop          float32   `db:"gdop"`
	Pdop          float32   `db:"pdop"`
	Hdop          float32   `db:"hdop"`
	Vdop          float32   `db:"vdop"`
	Tdop          float32   `db:"tdop"`
}
