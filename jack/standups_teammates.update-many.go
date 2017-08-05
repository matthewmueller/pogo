package jack

// GENERATED BY POGO. DO NOT EDIT.

import (
	"strings"
)

// UpdateMany rows by the condition
func (st *StandupsTeammates) UpdateMany(stst *StandupsTeammate, condition string, params ...interface{}) (standupsteammates []StandupsTeammate, err error) {
	fields := st.getFields(stst)

	// prepare the slices
	_c, _i, _v := querySlices(fields, len(params))

	// sql query
	sqlstr := `UPDATE jack.standups_teammates SET (` +
		strings.Join(_c, ", ") + `) = (` +
		strings.Join(_i, ", ") + `) ` +
		`WHERE ` + condition + ` ` +
		`RETURNING "standup_id", "teammate_id", "team_owner", "created_at", "updated_at"`

	values := []interface{}{}
	values = append(values, params...)
	values = append(values, _v...)

	// run query
	DBLog(sqlstr, values...)
	rows, err := st.DB.Query(sqlstr, values...)
	if err != nil {
		return standupsteammates, err
	}
	defer rows.Close()

	for rows.Next() {
		standupsteammate := StandupsTeammate{}
		err = rows.Scan(&standupsteammate.StandupID, &standupsteammate.TeammateID, &standupsteammate.TeamOwner, &standupsteammate.CreatedAt, &standupsteammate.UpdatedAt)
		if err != nil {
			return standupsteammates, err
		}
		standupsteammates = append(standupsteammates, standupsteammate)
	}

	if rows.Err() != nil {
		return standupsteammates, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(standupsteammates) == 0 {
		return make([]StandupsTeammate, 0), nil
	}

	return standupsteammates, nil
}