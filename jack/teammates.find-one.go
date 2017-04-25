package jack

import "github.com/matthewmueller/pgx"

// GENERATED BY POGO. DO NOT EDIT.

// FindOne Teammate by a condition
func (t *Teammates) FindOne(condition string, params ...interface{}) (teammate Teammate, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "slack_id", "team_id", "created", "updated"
    FROM public.teammates
    WHERE ` + condition

	DBLog(sqlstr, params...)
	row := t.DB.QueryRow(sqlstr, params...)
	err = row.Scan(&teammate.ID, &teammate.SlackID, &teammate.TeamID, &teammate.Created, &teammate.Updated)
	if err != nil {
		if err == pgx.ErrNoRows {
			return teammate, ErrTeammateNotFound
		}
		return teammate, err
	}

	return teammate, nil
}