// Include models that involved in business flow
package entities

import "database/sql"

type GemHistory struct {
	Base
	UserId      string         `db:"user_id"`
	SourceId    string         `db:"source_id"`
	Gems        float64        `db:"gems"`
	Type        string         `db:"type"`
	Status      string         `db:"status"`
	Reason      sql.NullString `db:"reason"`
	Metadata    sql.NullString `db:"metadata"`
	CollectedAt sql.NullTime   `db:"collected_at"`
}
