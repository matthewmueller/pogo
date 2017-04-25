package jack

import "strings"

// GENERATED BY POGO. DO NOT EDIT.

// Insert the TeammateStandup to the database.
func (ts *TeammateStandups) Insert(tsts *TeammateStandup) (teammatestandup TeammateStandup, err error) {
	// get all the non-nil fields and prepare them for the query
	c, i, v := querySlices(ts.getFields(tsts), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO public.teammate_standups (` + strings.Join(c, ", ") + `)
	VALUES (` + strings.Join(i, ", ") + `)
	RETURNING "id", "standup_id", "teammate_id", "time", "status"`

	DBLog(sqlstr, v...)
	row := ts.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&teammatestandup.ID, &teammatestandup.StandupID, &teammatestandup.TeammateID, &teammatestandup.Time, &teammatestandup.Status)
	if err != nil {
		return teammatestandup, err
	}

	return teammatestandup, nil
}