package tempo

// GENERATED BY POGO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"

	"github.com/matthewmueller/pgx"
)

// TaskStatus is the 'task_status' enum type from schema 'public'.
type TaskStatus uint16

const (

	// TaskStatusActive is the 'ACTIVE' TaskStatus.
	TaskStatusActive = TaskStatus(1)

	// TaskStatusPaused is the 'PAUSED' TaskStatus.
	TaskStatusPaused = TaskStatus(2)

	// TaskStatusStopped is the 'STOPPED' TaskStatus.
	TaskStatusStopped = TaskStatus(3)
)

// String returns the string value of the TaskStatus.
func (ts TaskStatus) String() string {
	var enumVal string

	switch ts {

	case TaskStatusActive:
		enumVal = "ACTIVE"

	case TaskStatusPaused:
		enumVal = "PAUSED"

	case TaskStatusStopped:
		enumVal = "STOPPED"

	}

	return enumVal
}

// MarshalText marshals TaskStatus into text.
func (ts TaskStatus) MarshalText() ([]byte, error) {
	return []byte(ts.String()), nil
}

// UnmarshalText unmarshals TaskStatus from text.
func (ts *TaskStatus) UnmarshalText(text []byte) error {
	switch string(text) {

	case "ACTIVE":
		*ts = TaskStatusActive

	case "PAUSED":
		*ts = TaskStatusPaused

	case "STOPPED":
		*ts = TaskStatusStopped

	default:
		return errors.New("invalid TaskStatus")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for TaskStatus.
func (ts TaskStatus) Value() (driver.Value, error) {
	return ts.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for TaskStatus.
func (ts *TaskStatus) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid TaskStatus")
	}

	return ts.UnmarshalText(buf)
}

// ScanPgx into PGX
func (ts *TaskStatus) ScanPgx(vr *pgx.ValueReader) error {
	if vr.Len() == -1 {
		return nil
	}
	return ts.UnmarshalText(vr.ReadBytes(vr.Len()))
}
