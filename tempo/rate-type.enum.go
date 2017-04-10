package tempo

// GENERATED BY POGO. DO NOT EDIT.

import (
 "database/sql/driver"
 "errors"
)

// RateType is the 'rate_type' enum type from schema 'public'.
type RateType uint16

const (

 // RateTypeCron is the 'CRON' RateType.
 RateTypeCron = RateType(1)

 // RateTypeRrule is the 'RRULE' RateType.
 RateTypeRrule = RateType(2)

 // RateTypeNlp is the 'NLP' RateType.
 RateTypeNlp = RateType(3)
)

// String returns the string value of the RateType.
func (rt RateType) String() string {
 var enumVal string

 switch rt {

 case RateTypeCron:
  enumVal = "CRON"

 case RateTypeRrule:
  enumVal = "RRULE"

 case RateTypeNlp:
  enumVal = "NLP"

 }

 return enumVal
}

// MarshalText marshals RateType into text.
func (rt RateType) MarshalText() ([]byte, error) {
 return []byte(rt.String()), nil
}

// UnmarshalText unmarshals RateType from text.
func (rt *RateType) UnmarshalText(text []byte) error {
 switch string(text) {

 case "CRON":
  *rt = RateTypeCron

 case "RRULE":
  *rt = RateTypeRrule

 case "NLP":
  *rt = RateTypeNlp

 default:
  return errors.New("invalid RateType")
 }

 return nil
}

// Value satisfies the sql/driver.Valuer interface for RateType.
func (rt RateType) Value() (driver.Value, error) {
 return rt.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for RateType.
func (rt *RateType) Scan(src interface{}) error {
 buf, ok := src.([]byte)
 if !ok {
  return errors.New("invalid RateType")
 }

 return rt.UnmarshalText(buf)
}
