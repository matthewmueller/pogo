package tempo

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrActiveNotFound returned if the Active is not found
var ErrActiveNotFound = errors.New("Active not found")

// Actives class
type Actives struct {
	DB DB
}

// Active model
type Active struct {
	ID        *uuid.UUID `json:"id,omitempty"`
	Token     *uuid.UUID `json:"token,omitempty"`
	State     *uuid.UUID `json:"state,omitempty"`
	Used      *bool      `json:"used,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewActive model
func NewActive(db DB) Actives {
	return Actives{
		DB: db,
	}
}

// getFields fetch the non-nil fields
func (a *Actives) getFields(aa *Active) map[string]interface{} {
	fields := map[string]interface{}{}

	if aa.ID != nil {
		fields["id"] = aa.ID
	}

	if aa.Token != nil {
		fields["token"] = aa.Token
	}

	if aa.State != nil {
		fields["state"] = aa.State
	}

	if aa.Used != nil {
		fields["used"] = aa.Used
	}

	if aa.CreatedAt != nil {
		fields["created_at"] = aa.CreatedAt
	}

	if aa.UpdatedAt != nil {
		fields["updated_at"] = aa.UpdatedAt
	}

	return fields
}