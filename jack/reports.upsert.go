package jack

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"

	"github.com/matthewmueller/pgx"
)

// Upsert the Report by the Primary Key
func (r *Reports) Upsert(rr *Report, action string) (report Report, err error) {
	fields := r.getFields(rr)

	// prepare the slices
	_c, _i, _v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(_c, ", ") + `) = ( EXCLUDED.` + strings.Join(_c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return report, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO jack.reports (` + strings.Join(_c, ", ") + `) ` +
		`VALUES (` + strings.Join(_i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "user_id", "timestamp", "questions", "standup_id", "status", "created_at", "updated_at"`

		// run query
	DBLog(sqlstr, _v...)
	row := r.DB.QueryRow(sqlstr, _v...)
	err = row.Scan(&report.ID, &report.UserID, &report.Timestamp, &report.Questions, &report.StandupID, &report.Status, &report.CreatedAt, &report.UpdatedAt)
	if err != nil && err != pgx.ErrNoRows {
		return report, err
	}

	return report, nil
}
