package jack

import (
	"errors"
	"time"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrTeamNotFound returned if the Team is not found
var ErrTeamNotFound = errors.New("Team not found")

// Teams class
type Teams struct {
	DB DB
}

// Team model
type Team struct {
	ID              *int       `json:"id,omitempty"`
	TeamAccessToken *string    `json:"team_access_token,omitempty"`
	BotAccessToken  *string    `json:"bot_access_token,omitempty"`
	TrialEnds       *time.Time `json:"trial_ends,omitempty"`
	TeamName        *string    `json:"team_name,omitempty"`
	TeamID          *string    `json:"team_id,omitempty"`
	BotUserID       *string    `json:"bot_user_id,omitempty"`
	Scope           *[]string  `json:"scope,omitempty"`
	Email           *string    `json:"email,omitempty"`
	StripeID        *string    `json:"stripe_id,omitempty"`
	Active          *bool      `json:"active,omitempty"`
	Created         *time.Time `json:"created,omitempty"`
	Updated         *time.Time `json:"updated,omitempty"`
	FreeTeammates   *int       `json:"free_teammates,omitempty"`
	CostPerUser     *int       `json:"cost_per_user,omitempty"`
}

// NewTeam model
func NewTeam(db DB) Teams {
	return Teams{
		DB: db,
	}
}

// getFields fetch the non-nil fields
func (t *Teams) getFields(tt *Team) map[string]interface{} {
	fields := map[string]interface{}{}

	if tt.ID != nil {
		fields["id"] = tt.ID
	}

	if tt.TeamAccessToken != nil {
		fields["team_access_token"] = tt.TeamAccessToken
	}

	if tt.BotAccessToken != nil {
		fields["bot_access_token"] = tt.BotAccessToken
	}

	if tt.TrialEnds != nil {
		fields["trial_ends"] = tt.TrialEnds
	}

	if tt.TeamName != nil {
		fields["team_name"] = tt.TeamName
	}

	if tt.TeamID != nil {
		fields["team_id"] = tt.TeamID
	}

	if tt.BotUserID != nil {
		fields["bot_user_id"] = tt.BotUserID
	}

	if tt.Scope != nil {
		fields["scope"] = tt.Scope
	}

	if tt.Email != nil {
		fields["email"] = tt.Email
	}

	if tt.StripeID != nil {
		fields["stripe_id"] = tt.StripeID
	}

	if tt.Active != nil {
		fields["active"] = tt.Active
	}

	if tt.Created != nil {
		fields["created"] = tt.Created
	}

	if tt.Updated != nil {
		fields["updated"] = tt.Updated
	}

	if tt.FreeTeammates != nil {
		fields["free_teammates"] = tt.FreeTeammates
	}

	if tt.CostPerUser != nil {
		fields["cost_per_user"] = tt.CostPerUser
	}

	return fields
}