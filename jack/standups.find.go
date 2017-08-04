package jack

import (
	"github.com/matthewmueller/pgx"
	"github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// Find a Standup by "id"
func (s *Standups) Find(id *uuid.UUID) (standup Standup, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
    FROM jack.standups
    WHERE "id" = $1`

	DBLog(sqlstr, id)
	row := s.DB.QueryRow(sqlstr, id)
	err = row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return standup, ErrStandupNotFound
		}
		return standup, err
	}

	return standup, nil
}

// FindBySlackChannelID find a Standup
func (s *Standups) FindBySlackChannelID(slackchannelid *string) (standup Standup, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
		SELECT "id", "name", "slack_channel_id", "time", "timezone", "questions", "team_id", "created_at", "updated_at"
		FROM jack.standups
		WHERE "slack_channel_id" = $1`

	DBLog(sqlstr, slackchannelid)
	row := s.DB.QueryRow(sqlstr, slackchannelid)
	err = row.Scan(&standup.ID, &standup.Name, &standup.SlackChannelID, &standup.Time, &standup.Timezone, &standup.Questions, &standup.TeamID, &standup.CreatedAt, &standup.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return standup, ErrStandupNotFound
		}
		return standup, err
	}

	return standup, nil
}
