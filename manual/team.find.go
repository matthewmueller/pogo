package model

// Find a team by id
func (m *Teams) Find(ID *string) (team Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT id, slack_team_id, slack_team_access_token, slack_bot_access_token, slack_bot_id, team_name, scope, email, stripe_id, active, free_teammates, cost_per_user, trial_ends, created_at, updated_at
    FROM jack.teams
    WHERE id = $1`

	DBLog(sqlstr, ID)
	row := m.DB.QueryRow(sqlstr, ID)
	err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		return team, err
	}

	return team, nil
}
