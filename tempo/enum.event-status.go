package tempo

// GENERATED BY POGO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"

	"github.com/matthewmueller/pgx"
)

// EventStatus is the 'event_status' enum type from schema 'public'.
type EventStatus uint16

const (

	// EventStatusReady is the 'READY' EventStatus.
	EventStatusReady = EventStatus(1)

	// EventStatusRunning is the 'RUNNING' EventStatus.
	EventStatusRunning = EventStatus(2)

	// EventStatusSuccess is the 'SUCCESS' EventStatus.
	EventStatusSuccess = EventStatus(3)

	// EventStatusFailure is the 'FAILURE' EventStatus.
	EventStatusFailure = EventStatus(4)

	// EventStatusSkipped is the 'SKIPPED' EventStatus.
	EventStatusSkipped = EventStatus(5)

	// EventStatusPaused is the 'PAUSED' EventStatus.
	EventStatusPaused = EventStatus(6)
)

// String returns the string value of the EventStatus.
func (es EventStatus) String() string {
	var enumVal string

	switch es {

	case EventStatusReady:
		enumVal = "READY"

	case EventStatusRunning:
		enumVal = "RUNNING"

	case EventStatusSuccess:
		enumVal = "SUCCESS"

	case EventStatusFailure:
		enumVal = "FAILURE"

	case EventStatusSkipped:
		enumVal = "SKIPPED"

	case EventStatusPaused:
		enumVal = "PAUSED"

	}

	return enumVal
}

// MarshalText marshals EventStatus into text.
func (es EventStatus) MarshalText() ([]byte, error) {
	return []byte(es.String()), nil
}

// UnmarshalText unmarshals EventStatus from text.
func (es *EventStatus) UnmarshalText(text []byte) error {
	switch string(text) {

	case "READY":
		*es = EventStatusReady

	case "RUNNING":
		*es = EventStatusRunning

	case "SUCCESS":
		*es = EventStatusSuccess

	case "FAILURE":
		*es = EventStatusFailure

	case "SKIPPED":
		*es = EventStatusSkipped

	case "PAUSED":
		*es = EventStatusPaused

	default:
		return errors.New("invalid EventStatus")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for EventStatus.
func (es EventStatus) Value() (driver.Value, error) {
	return es.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for EventStatus.
func (es *EventStatus) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid EventStatus")
	}

	return es.UnmarshalText(buf)
}

// ScanPgx into PGX
func (es *EventStatus) ScanPgx(vr *pgx.ValueReader) error {
	if vr.Len() == -1 {
		return nil
	}
	return es.UnmarshalText(vr.ReadBytes(vr.Len()))
}
