package jack

import "github.com/matthewmueller/pgx"

// GENERATED BY POGO. DO NOT EDIT.

// FindOne Teammate by a condition
func (t *Teammates) FindOne(condition string, params ...interface{}) (teammate Teammate, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
    FROM jack.teammates
    WHERE ` + condition

	DBLog(sqlstr, params...)
	row := t.DB.QueryRow(sqlstr, params...)
	err = row.Scan(&teammate.ID, &teammate.SlackID, &teammate.Username, &teammate.FirstName, &teammate.LastName, &teammate.Email, &teammate.Avatar, &teammate.Timezone, &teammate.CreatedAt, &teammate.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return teammate, ErrTeammateNotFound
		}
		return teammate, err
	}

	return teammate, nil
}
