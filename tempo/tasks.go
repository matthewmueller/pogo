package tempo

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

// GENERATED BY POGO. DO NOT EDIT.

// ErrTaskNotFound returned if the Task is not found
var ErrTaskNotFound = errors.New("Task not found")

// Tasks class
type Tasks struct {
	DB DB
}

// Task model
type Task struct {
	ID            *uuid.UUID              `json:"id,omitempty"`
	Key           *string                 `json:"key,omitempty"`
	Target        *string                 `json:"target,omitempty"`
	Rate          *string                 `json:"rate,omitempty"`
	Offset        *time.Time              `json:"offset,omitempty"`
	Timezone      *string                 `json:"timezone,omitempty"`
	RateType      *TaskRateType           `json:"rate_type,omitempty"`
	RateOptions   *map[string]interface{} `json:"rate_options,omitempty"`
	Status        *TaskStatus             `json:"status,omitempty"`
	TargetType    *TaskTargetType         `json:"target_type,omitempty"`
	TargetOptions *map[string]interface{} `json:"target_options,omitempty"`
	User          *uuid.UUID              `json:"user,omitempty"`
	Meta          *map[string]interface{} `json:"meta,omitempty"`
	RefreshedAt   *time.Time              `json:"refreshed_at,omitempty"`
	CreatedAt     *time.Time              `json:"created_at,omitempty"`
	UpdatedAt     *time.Time              `json:"updated_at,omitempty"`
}

// NewTask model
func NewTask(db DB) Tasks {
	return Tasks{
		DB: db,
	}
}

// getFields fetch the non-nil fields
func (t *Tasks) getFields(tt *Task) map[string]interface{} {
	fields := map[string]interface{}{}

	if tt.ID != nil {
		fields["id"] = tt.ID
	}

	if tt.Key != nil {
		fields["key"] = tt.Key
	}

	if tt.Target != nil {
		fields["target"] = tt.Target
	}

	if tt.Rate != nil {
		fields["rate"] = tt.Rate
	}

	if tt.Offset != nil {
		fields["offset"] = tt.Offset
	}

	if tt.Timezone != nil {
		fields["timezone"] = tt.Timezone
	}

	if tt.RateType != nil {
		fields["rate_type"] = tt.RateType
	}

	if tt.RateOptions != nil {
		fields["rate_options"] = tt.RateOptions
	}

	if tt.Status != nil {
		fields["status"] = tt.Status
	}

	if tt.TargetType != nil {
		fields["target_type"] = tt.TargetType
	}

	if tt.TargetOptions != nil {
		fields["target_options"] = tt.TargetOptions
	}

	if tt.User != nil {
		fields["user"] = tt.User
	}

	if tt.Meta != nil {
		fields["meta"] = tt.Meta
	}

	if tt.RefreshedAt != nil {
		fields["refreshed_at"] = tt.RefreshedAt
	}

	if tt.CreatedAt != nil {
		fields["created_at"] = tt.CreatedAt
	}

	if tt.UpdatedAt != nil {
		fields["updated_at"] = tt.UpdatedAt
	}

	return fields
}
