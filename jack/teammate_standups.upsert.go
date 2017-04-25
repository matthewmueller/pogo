package jack

// GENERATED BY POGO. DO NOT EDIT.

import (
	"errors"
	"strings"

	"github.com/matthewmueller/pgx"
)

// Upsert the TeammateStandup by the Primary Key
func (ts *TeammateStandups) Upsert(tsts *TeammateStandup, action string) (teammatestandup TeammateStandup, err error) {
	fields := ts.getFields(tsts)

	// prepare the slices
	c, i, v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return teammatestandup, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO public.teammate_standups (` + strings.Join(c, ", ") + `) ` +
		`VALUES (` + strings.Join(i, ", ") + `) ` +
		`ON CONFLICT ("id") ` +
		upsertAction + ` ` +
		`RETURNING "id", "standup_id", "teammate_id", "time", "status"`

		// run query
	DBLog(sqlstr, v...)
	row := ts.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&teammatestandup.ID, &teammatestandup.StandupID, &teammatestandup.TeammateID, &teammatestandup.Time, &teammatestandup.Status)
	if err != nil && err != pgx.ErrNoRows {
		return teammatestandup, err
	}

	return teammatestandup, nil
}

// UpsertByStandupIDAndTeammateID find a TeammateStandup
func (ts *TeammateStandups) UpsertByStandupIDAndTeammateID(tsts *TeammateStandup, action string) (teammatestandup TeammateStandup, err error) {
	fields := ts.getFields(tsts)

	// prepare the slices
	c, i, v := querySlices(fields, 0)

	// determine on conflict action
	var upsertAction string
	if action == UpsertDoUpdate {
		upsertAction = `DO UPDATE SET (` + strings.Join(c, ", ") + `) = ( EXCLUDED.` + strings.Join(c, ", EXCLUDED.") + `)`
	} else if action == UpsertDoNothing {
		upsertAction = UpsertDoNothing
	} else {
		return teammatestandup, errors.New("invalid upsert action")
	}

	// sql query
	sqlstr := `INSERT INTO public.teammate_standups (` + strings.Join(c, ", ") + `) ` +
		`VALUES (` + strings.Join(i, ", ") + `) ` +
		`ON CONFLICT ("standupid", "teammateid") ` +
		upsertAction + ` ` +
		`RETURNING "id", "standup_id", "teammate_id", "time", "status"`

		// run query
	DBLog(sqlstr, v...)
	row := ts.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&teammatestandup.ID, &teammatestandup.StandupID, &teammatestandup.TeammateID, &teammatestandup.Time, &teammatestandup.Status)
	if err != nil && err != pgx.ErrNoRows {
		return teammatestandup, err
	}

	return teammatestandup, nil
}
