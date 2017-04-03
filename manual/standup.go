package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Standup type
type Standup struct {
	ID             uuid.NullUUID           `json:"id"`
	Name           *string                 `json:"name"`
	SlackChannelID *string                 `json:"slack_channel_id"`
	Time           *string                 `json:"time"`
	Timezone       *string                 `json:"timezone"`
	Questions      *map[string]interface{} `json:"questions"`
	TeamID         *uuid.UUID              `json:"team_id"`
	CreatedAt      *time.Time              `json:"created_at"`
	UpdatedAt      *time.Time              `json:"updated_at"`
}

// GetFields the non-nil fields
func (s *Standup) getFields() map[string]interface{} {
	fields := map[string]interface{}{}

	if s.ID.Valid {
		fields["id"] = s.ID.UUID
	}
	if s.Name != nil {
		fields["name"] = s.Name
	}
	if s.SlackChannelID != nil {
		fields["slack_channel_id"] = s.SlackChannelID
	}
	if s.Time != nil {
		fields["time"] = s.Time
	}
	if s.Timezone != nil {
		fields["timezone"] = s.Timezone
	}
	if s.Questions != nil {
		fields["questions"] = s.Questions
	}
	if s.TeamID != nil {
		fields["team_id"] = s.TeamID
	}
	if s.CreatedAt != nil {
		fields["created_at"] = s.CreatedAt
	}
	if s.UpdatedAt != nil {
		fields["updated_at"] = s.UpdatedAt
	}
	return fields
}
