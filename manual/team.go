package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Team type
type Team struct {
	ID                   *uuid.UUID `json:"id,omitempty"`                      // id
	SlackTeamID          *string    `json:"slack_team_id,omitempty"`           // slack_team_id
	SlackTeamAccessToken *string    `json:"slack_team_access_token,omitempty"` // slack_team_access_token
	SlackBotAccessToken  *string    `json:"slack_bot_access_token,omitempty"`  // slack_bot_access_token
	SlackBotID           *string    `json:"slack_bot_id,omitempty"`            // slack_bot_id
	TeamName             *string    `json:"team_name,omitempty"`               // team_name
	Scope                []*string  `json:"scope,omitempty"`                   // scope
	Email                *string    `json:"email,omitempty"`                   // email
	StripeID             *string    `json:"stripe_id,omitempty"`               // stripe_id
	Active               *bool      `json:"active,omitempty"`                  // active
	FreeTeammates        *int       `json:"free_teammates,omitempty"`          // free_teammates
	CostPerUser          *int       `json:"cost_per_user,omitempty"`           // cost_per_user
	TrialEnds            *time.Time `json:"trial_ends,omitempty"`              // trial_ends
	CreatedAt            *time.Time `json:"created_at,omitempty"`              // created_at
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`              // updated_at
}

// GetFields the non-nil fields
func (t *Team) getFields() map[string]interface{} {
	fields := map[string]interface{}{}

	if t.ID != nil {
		fields["id"] = t.ID
	}
	if t.SlackTeamID != nil {
		fields["slack_team_id"] = t.SlackTeamID
	}
	if t.SlackTeamAccessToken != nil {
		fields["slack_team_access_token"] = t.SlackTeamAccessToken
	}
	if t.SlackBotAccessToken != nil {
		fields["slack_bot_access_token"] = t.SlackBotAccessToken
	}
	if t.SlackBotID != nil {
		fields["slack_bot_id"] = t.SlackBotID
	}
	if t.TeamName != nil {
		fields["team_name"] = t.TeamName
	}
	if t.Scope != nil {
		fields["scope"] = t.Scope
	}
	if t.Email != nil {
		fields["email"] = t.Email
	}
	if t.StripeID != nil {
		fields["stripe_id"] = t.StripeID
	}
	if t.Active != nil {
		fields["active"] = t.Active
	}
	if t.FreeTeammates != nil {
		fields["free_teammates"] = t.FreeTeammates
	}
	if t.CostPerUser != nil {
		fields["cost_per_user"] = t.CostPerUser
	}
	if t.TrialEnds != nil {
		fields["trial_ends"] = t.TrialEnds
	}
	if t.CreatedAt != nil {
		fields["created_at"] = t.CreatedAt
	}
	if t.UpdatedAt != nil {
		fields["updated_at"] = t.UpdatedAt
	}

	return fields
}
