package jack

import "github.com/matthewmueller/pgx"

// GENERATED BY POGO. DO NOT EDIT.

// FindOne Standup by a condition
func (s *Standups) FindOne(condition string, params ...interface{}) (standup Standup, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "channel_id", "time", "tz", "questions", "created", "updated", "owner", "team_id", "name"
    FROM public.standups
    WHERE ` + condition

	DBLog(sqlstr, params...)
	row := s.DB.QueryRow(sqlstr, params...)
	err = row.Scan(&standup.ID, &standup.ChannelID, &standup.Time, &standup.Tz, &standup.Questions, &standup.Created, &standup.Updated, &standup.Owner, &standup.TeamID, &standup.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return standup, ErrStandupNotFound
		}
		return standup, err
	}

	return standup, nil
}