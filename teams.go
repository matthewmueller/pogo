package jack

import (
	"time"

	"github.com/jackc/pgx"
	uuid "github.com/satori/go.uuid"
)

// Teams class
type Teams struct {
	db *pgx.Conn
}

// Team model
type Team struct {
	ID                   *uuid.UUID `json:"id,omitempty"`
	SlackTeamID          *string    `json:"slack_team_id,omitempty"`
	SlackTeamAccessToken *string    `json:"slack_team_access_token,omitempty"`
	SlackBotAccessToken  *string    `json:"slack_bot_access_token,omitempty"`
	SlackBotID           *string    `json:"slack_bot_id,omitempty"`
	TeamName             *string    `json:"team_name,omitempty"`
	Scope                *[]string  `json:"scope,omitempty"`
	Email                *string    `json:"email,omitempty"`     // user email
	StripeID             *string    `json:"stripe_id,omitempty"` // user stripe id
	Active               *bool      `json:"active,omitempty"`
	FreeTeammates        *int       `json:"free_teammates,omitempty"`
	CostPerUser          *int       `json:"cost_per_user,omitempty"`
	TrialEnds            *time.Time `json:"trial_ends,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
}

// team constructor
func team(db *pgx.Conn) *Teams {
	return &Teams{db}
}

// get all the non-nil fields
func fields(team *Team) map[string]interface{} {
	fields := make(map[string]interface{})

	if team.ID != nil {
		fields["id"] = team.ID
	}
	if team.SlackTeamID != nil {
		fields["slack_team_id"] = team.SlackTeamID
	}
	if team.SlackTeamAccessToken != nil {
		fields["slack_team_access_token"] = team.SlackTeamAccessToken
	}
	if team.SlackBotAccessToken != nil {
		fields["slack_bot_access_token"] = team.SlackBotAccessToken
	}
	if team.SlackBotID != nil {
		fields["slack_bot_id"] = team.SlackBotID
	}
	if team.TeamName != nil {
		fields["team_name"] = team.TeamName
	}
	if team.Scope != nil {
		fields["scope"] = team.Scope
	}
	if team.Email != nil {
		fields["email"] = team.Email
	}
	if team.StripeID != nil {
		fields["stripe_id"] = team.StripeID
	}
	if team.Active != nil {
		fields["active"] = team.Active
	}
	if team.FreeTeammates != nil {
		fields["free_teammates"] = team.FreeTeammates
	}
	if team.CostPerUser != nil {
		fields["cost_per_user"] = team.CostPerUser
	}
	if team.TrialEnds != nil {
		fields["trial_ends"] = team.TrialEnds
	}
	if team.CreatedAt != nil {
		fields["created_at"] = team.CreatedAt
	}
	if team.UpdatedAt != nil {
		fields["updated_at"] = team.UpdatedAt
	}

	return fields
}
