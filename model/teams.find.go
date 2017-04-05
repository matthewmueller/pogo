package model

// GENERATED BY POGO. DO NOT EDIT.

// Find a team by id
func (t *Teams) Find(id *string) (team Team, err error) {
 // sql select query, primary key provided by sequence
 sqlstr := `
    SELECT id, slack_team_id, slack_team_access_token, slack_bot_access_token, slack_bot_id, team_name, scope, email, stripe_id, active, free_teammates, cost_per_user, trial_ends, created_at, updated_at
    FROM jack.teams
    WHERE id = $1`

 XOLog(sqlstr, id)
 row := t.DB.QueryRow(sqlstr, id)
 err = row.Scan(&team.ID, &team.SlackTeamID, &team.SlackTeamAccessToken, &team.SlackBotAccessToken, &team.SlackBotID, &team.TeamName, &team.Scope, &team.Email, &team.StripeID, &team.Active, &team.FreeTeammates, &team.CostPerUser, &team.TrialEnds, &team.CreatedAt, &team.UpdatedAt)
 if err != nil {
  return team, err
 }

 return team, nil
}
