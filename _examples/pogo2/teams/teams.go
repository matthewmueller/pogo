package teams

import (
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/_examples/pogo2/pogo"
	"github.com/matthewmueller/pogo/_examples/pogo2/team"
	uuid "github.com/satori/go.uuid"
)

// Teams table
type Teams struct {
	db pogo.DB
}

// columns on the team
type columns struct {
	ID          *uuid.UUID
	SlackTeamID *string
}

// Team struct
type Team struct {
	columns *columns
}

// New creates a team
func New() *Team {
	return &Team{
		columns: &columns{},
	}
}

func (c *Team) ID(id uuid.UUID) *Team {
	c.columns.ID = &id
	return c
}

func (c *Team) GetID() *uuid.UUID {
	return c.columns.ID
}

func (teams *Teams) Find(db pogo.DB, id string) (team *team.Team, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
	SELECT "id", "slack_team_id", "slack_team_access_token", "slack_bot_access_token", "slack_bot_id", "team_name", "scope", "email", "stripe_id", "active", "free_teammates", "cost_per_user", "trial_ends", "created_at", "updated_at"
	FROM jack.teams
	WHERE "id" = $1
	`

	Log(sqlstr, id)
	row := teams.db.QueryRow(sqlstr, id)
	// tm := team.
	// 	ID(123).
	// 	SlackTeamID("123")

	if e := row.Scan(team.ID, team.SlackTeamID, team.SlackTeamAccessToken, team.SlackBotAccessToken, team.SlackBotID, team.TeamName, team.Scope, team.Email, team.StripeID, team.Active, team.FreeTeammates, team.CostPerUser, team.TrialEnds, team.CreatedAt, team.UpdatedAt); e != nil {
		if e == pgx.ErrNoRows {
			return nil, ErrTeamNotFound
		}
		return nil, e
	}

	return team, nil
}
