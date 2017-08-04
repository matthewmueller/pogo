package jack

// GENERATED BY POGO. DO NOT EDIT.

// FindMany find many Teammates by a condition
func (t *Teammates) FindMany(condition string, params ...interface{}) (teammates []Teammate, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "slack_id", "username", "first_name", "last_name", "email", "avatar", "timezone", "created_at", "updated_at"
    FROM jack.teammates
    WHERE ` + condition

	DBLog(sqlstr, params...)
	rows, err := t.DB.Query(sqlstr, params...)
	if err != nil {
		return teammates, err
	}
	defer rows.Close()

	for rows.Next() {
		teammate := Teammate{}
		err = rows.Scan(&teammate.ID, &teammate.SlackID, &teammate.Username, &teammate.FirstName, &teammate.LastName, &teammate.Email, &teammate.Avatar, &teammate.Timezone, &teammate.CreatedAt, &teammate.UpdatedAt)
		if err != nil {
			return teammates, err
		}
		teammates = append(teammates, teammate)
	}

	if rows.Err() != nil {
		return teammates, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(teammates) == 0 {
		return make([]Teammate, 0), nil
	}

	return teammates, nil
}
