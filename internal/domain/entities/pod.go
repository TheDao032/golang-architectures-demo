// Include models that involved in business flow
package entities

import (
	"time"

	"github.com/lib/pq"
)

type POD struct {
	Time          time.Time       `db:"time"`
	ExperimentId  int             `db:"experiment_id"`
	ApplicationId int             `db:"application_id"`
	Wn            int             `db:"wn"`
	Tow           int64           `db:"tow"`
	Decimals      float32         `db:"decimals"`
	NSat          int             `db:"n_sat"`
	PosX          float32         `db:"pos_x"`
	PosY          float32         `db:"pos_y"`
	PosZ          float32         `db:"pos_z"`
	VelX          float32         `db:"vel_x"`
	VelY          float32         `db:"vel_y"`
	VelZ          float32         `db:"vel_z"`
	PosStd        pq.Float32Array `db:"pos_std"`
	VelStd        pq.Float32Array `db:"vel_std"`
	ClockBias     float32         `db:"clock_bias"`
	ClockDrift    float32         `db:"clock_drift"`
	AmbigVec      pq.Float32Array `db:"ambig_vec"`
	AmbigAcc      pq.Float32Array `db:"ambig_acc"`
}
