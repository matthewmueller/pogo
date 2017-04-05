package model

import (
 "strings"

 "github.com/pkg/errors"
)

// GENERATED BY POGO. DO NOT EDIT.

// Insert the Standup to the database.
func (s *Standups) Insert(ss *Standup) (standup Standup, err error) {
 // get all the non-nil fields and prepare them for the query
 c, i, v := querySlices(s.getFields(ss), 0)

 // sql insert query, primary key provided by sequence
 sqlstr := `
	INSERT INTO jack.standups (` + strings.Join(c, ", ") + `)
	VALUES (` + strings.Join(i, ", ") + `)
	RETURNING id, name, slack_channel_id, time, timezone, questions, team_id, created_at, updated_at`

 XOLog(sqlstr, v...)
 row := s.DB.QueryRow(sqlstr, v...)
 err = row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt)
 if err != nil {
  return standup, errors.Wrap(err, "could not insert into 'standups'")
 }

 return standup, nil
}
