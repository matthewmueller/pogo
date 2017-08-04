package jack

import "strings"

// GENERATED BY POGO. DO NOT EDIT.

// Insert the Team to the database.
func (t *Teams) Insert(tt *Team) (team Team, err error) {
	// get all the non-nil fields and prepare them for the query
	_c, _i, _v := querySlices(t.getFields(tt), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	INSERT INTO jack.teams (` + strings.Join(_c, ", ") + `)
	VALUES (` + strings.Join(_i, ", ") + `)
	RETURNING "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"`

	DBLog(sqlstr, _v...)
	row := t.DB.QueryRow(sqlstr, _v...)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		return team, err
	}

	return team, nil
}
