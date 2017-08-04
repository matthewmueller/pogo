package jack

// GENERATED BY POGO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"

	"github.com/matthewmueller/pgx"
)

// ReportStatus is the 'report_status' enum type from schema 'jack'.
type ReportStatus uint16

const (

	// ReportStatusAsked is the 'ASKED' ReportStatus.
	ReportStatusAsked = ReportStatus(1)

	// ReportStatusSkip is the 'SKIP' ReportStatus.
	ReportStatusSkip = ReportStatus(2)

	// ReportStatusComplete is the 'COMPLETE' ReportStatus.
	ReportStatusComplete = ReportStatus(3)
)

// String returns the string value of the ReportStatus.
func (rs ReportStatus) String() string {
	var enumVal string

	switch rs {

	case ReportStatusAsked:
		enumVal = "ASKED"

	case ReportStatusSkip:
		enumVal = "SKIP"

	case ReportStatusComplete:
		enumVal = "COMPLETE"

	}

	return enumVal
}

// MarshalText marshals ReportStatus into text.
func (rs ReportStatus) MarshalText() ([]byte, error) {
	return []byte(rs.String()), nil
}

// UnmarshalText unmarshals ReportStatus from text.
func (rs *ReportStatus) UnmarshalText(text []byte) error {
	switch string(text) {

	case "ASKED":
		*rs = ReportStatusAsked

	case "SKIP":
		*rs = ReportStatusSkip

	case "COMPLETE":
		*rs = ReportStatusComplete

	default:
		return errors.New("invalid ReportStatus")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for ReportStatus.
func (rs ReportStatus) Value() (driver.Value, error) {
	return rs.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for ReportStatus.
func (rs *ReportStatus) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid ReportStatus")
	}

	return rs.UnmarshalText(buf)
}

// ScanPgx into PGX
func (rs *ReportStatus) ScanPgx(vr *pgx.ValueReader) error {
	if vr.Len() == -1 {
		return nil
	}
	return rs.UnmarshalText(vr.ReadBytes(vr.Len()))
}
