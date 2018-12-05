package schema

// Schema struct is a universal
// database structure for storing
// different types of database
// schemas
type Schema struct {
	Provider   string
	Name       string
	Tables     []*Table
	Enums      []*Enum
	Procedures []*Procedure
	Coerce     Coercer
}

// Coercer coerces SQL types into Go types
type Coercer interface {
	Type(sqlType string) (goType string, err error)
	FilterType(sqlType string) (filterType FilterType, err error)
}
