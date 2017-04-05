package model

import (
	"strings"
)

// Insert a standup
func (s *Standup) Insert(db DB) (standup Standup, err error) {
	// get all the non-nil fields and prepare them for the query
	c, i, v := querySlices(s.getFields(), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
		INSERT INTO
		jack.standups (` + strings.Join(c, ", ") + `)
		VALUES (` + strings.Join(i, ", ") + `)
		RETURNING id, name, slack_channel_id, time, timezone, questions, team_id, created_at, updated_at`

	DBLog(sqlstr, v...)
	row := db.QueryRow(sqlstr, v...)
	err = row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt)
	if err != nil {
		return standup, err
	}

	return standup, nil
}
