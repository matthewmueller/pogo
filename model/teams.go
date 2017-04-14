package model

import "time"

// GENERATED BY POGO. DO NOT EDIT.

// Teams class
type Teams struct {
 DB DB
}

// Team model
type Team struct {
 ID                   *string    `json:"id,omitempty"`
 SlackTeamID          *string    `json:"slack_team_id,omitempty"`
 SlackTeamAccessToken *string    `json:"slack_team_access_token,omitempty"`
 SlackBotAccessToken  *string    `json:"slack_bot_access_token,omitempty"`
 SlackBotID           *string    `json:"slack_bot_id,omitempty"`
 TeamName             *string    `json:"team_name,omitempty"`
 Scope                *[]string  `json:"scope,omitempty"`
 Email                *string    `json:"email,omitempty"`
 StripeID             *string    `json:"stripe_id,omitempty"`
 Active               *bool      `json:"active,omitempty"`
 FreeTeammates        *int       `json:"free_teammates,omitempty"`
 CostPerUser          *int       `json:"cost_per_user,omitempty"`
 TrialEnds            *time.Time `json:"trial_ends,omitempty"`
 CreatedAt            *time.Time `json:"created_at,omitempty"`
 UpdatedAt            *time.Time `json:"updated_at,omitempty"`
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

 if tt.SlackTeamID != nil {
  fields["slack_team_id"] = tt.SlackTeamID
 }

 if tt.SlackTeamAccessToken != nil {
  fields["slack_team_access_token"] = tt.SlackTeamAccessToken
 }

 if tt.SlackBotAccessToken != nil {
  fields["slack_bot_access_token"] = tt.SlackBotAccessToken
 }

 if tt.SlackBotID != nil {
  fields["slack_bot_id"] = tt.SlackBotID
 }

 if tt.TeamName != nil {
  fields["team_name"] = tt.TeamName
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

 if tt.FreeTeammates != nil {
  fields["free_teammates"] = tt.FreeTeammates
 }

 if tt.CostPerUser != nil {
  fields["cost_per_user"] = tt.CostPerUser
 }

 if tt.TrialEnds != nil {
  fields["trial_ends"] = tt.TrialEnds
 }

 if tt.CreatedAt != nil {
  fields["created_at"] = tt.CreatedAt
 }

 if tt.UpdatedAt != nil {
  fields["updated_at"] = tt.UpdatedAt
 }

 return fields
}