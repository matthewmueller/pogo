package schema

// New creates a new schema
func New(
	provider string,
	name string,
	tables []*Table,
	enums []*Enum,
	procedures []*Procedure,
) *Schema {
	return &Schema{
		provider,
		name,
		tables,
		enums,
		procedures,
	}
}

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
