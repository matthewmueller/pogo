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
}
