package model

import (
	"strings"
)

// Insert inserts the Team to the database.
func (m *Teams) Insert(t *Team) (team Team, err error) {
	// get all the non-nil fields and prepare them for the query
	c, i, v := querySlices(m.getFields(t), 0)

	// sql insert query, primary key provided by sequence
	sqlstr := `
	  INSERT INTO
	  jack.teams (` + strings.Join(c, ", ") + `)
	  VALUES (` + strings.Join(i, ", ") + `)
	  RETURNING id, slack_team_id, slack_team_access_token, slack_bot_access_token, slack_bot_id, team_name, scope, email, stripe_id, active, free_teammates, cost_per_user, trial_ends, created_at, updated_at`

	XOLog(sqlstr, v...)
	row := m.DB.QueryRow(sqlstr, v...)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		return team, err
	}

	return team, nil
}
