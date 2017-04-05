package model

import (
	"errors"
	"strings"
)

// Update the Team by the Primary Key
func (m *Teams) Update(id *string, t *Team) (team Team, err error) {
	fields := m.getFields(t)

	// first check if we have the primary key
	if id == nil {
		return team, errors.New(`primary key "id" must be non-nil`)
	}

	// don't update the primary key
	// if it's present in t
	delete(fields, "id")

	// prepare the slices
	c, i, v := querySlices(fields, 1)

	// sql query
	sqlstr := `UPDATE jack.teams SET (` +
		strings.Join(c, ", ") + `) = (` +
		strings.Join(i, ", ") + `)
		WHERE id = $1
		RETURNING id, slack_team_id, slack_team_access_token, slack_bot_access_token, slack_bot_id, team_name, scope, email, stripe_id, active, free_teammates, cost_per_user, trial_ends, created_at, updated_at`

	// run query
	values := append([]interface{}{t.ID}, v...)
	DBLog(sqlstr, values...)

	row := m.DB.QueryRow(sqlstr, values...)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		return team, err
	}

	return team, nil
}
