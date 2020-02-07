package schema

import (
	"fmt"
	"strings"
)

// Schema struct
type Schema struct {
	Name       string
	Tables     Tables
	Views      []*View
	Enums      Enums
	Procedures []*Procedure
}

const br = "\n\n"

func (s *Schema) String() string {
	return strings.TrimSpace(
		s.Enums.String() + br + s.Tables.String(),
	)
}

// Enums list
type Enums []*Enum

func (enums Enums) String() string {
	blocks := make([]string, len(enums))
	for i, enum := range enums {
		blocks[i] = enum.String()
	}
	return strings.Join(blocks, "\n")
}

// Tables list
type Tables []*Table

func (tables Tables) String() string {
	blocks := make([]string, len(tables))
	for i, table := range tables {
		blocks[i] = table.String()
	}
	return strings.Join(blocks, "\n\n")
}

// Table struct
type Table struct {
	Schema  string
	Name    string
	Columns Columns

	// constraints
	Primary  *Primary
	Foreigns []*Foreign
	Uniques  []*Unique
}

func (t *Table) String() string {
	return fmt.Sprintf("create table %q.%q (\n\t%s,\n\t%s\n);", t.Schema, t.Name, t.Columns.Join(",\n\t"), t.Primary.String())
}

// DataType type
type DataType string

// String
func (datatype DataType) String() string {
	return string(datatype)
}

// Columns list
type Columns []*Column

// Join columns by a separator
func (columns Columns) Join(sep string) string {
	blocks := make([]string, len(columns))
	for i, column := range columns {
		blocks[i] = column.String()
	}
	return strings.Join(blocks, sep)
}

// Column struct
type Column struct {
	Order   int
	Name    string
	Type    DataType
	NotNull bool
	Comment *string
	Default *string
}

func (c *Column) String() string {
	var s strings.Builder
	s.WriteString(c.Name + " " + c.Type.String())
	if c.NotNull {
		s.WriteString(" not null")
	}
	if c.Default != nil {
		s.WriteString(" default " + *c.Default)
	}
	// TODO: handle comment
	return s.String()
}

// Enum data type
type Enum struct {
	Schema string
	Name   string
	Values []*EnumValue
}

func (e *Enum) String() string {
	values := make([]string, len(e.Values))
	for i, value := range e.Values {
		values[i] = value.String()
	}
	return fmt.Sprintf(`create type %q.%q as enum (%s)`, e.Schema, e.Name, strings.Join(values, ", "))
}

// EnumValue is an enum value
type EnumValue struct {
	Order int
	Label string
}

func (e *EnumValue) String() string {
	return `'` + e.Label + `'`
}

// Primary constraint
type Primary struct {
	Schema  string
	Name    string
	Columns Columns
}

func (p *Primary) String() string {
	columns := make([]string, len(p.Columns))
	for i, column := range p.Columns {
		columns[i] = column.Name
	}
	return fmt.Sprintf(`constraint %q primary key (%q)`, p.Name, strings.Join(columns, ", "))
}

// Foreign constraint
type Foreign struct {
	Schema     string
	Name       string
	Columns    Columns
	RefTable   string
	RefColumns Columns
}

func (f *Foreign) String() string {
	return fmt.Sprintf(`constraint %q foreign key (%q) references %q.%q (%q)`, f.Name, f.Columns.Join(", "), f.Schema, f.RefTable, f.RefColumns.Join(", "))
}

// Unique constraint
type Unique struct {
	Schema  string
	Name    string
	Columns []*Column
}

func (u *Unique) String() string {
	return ``
}

// Procedure represents a stored procedure.
type Procedure struct {
	Name       string // proc name
	Params     []*ProcedureParam
	ReturnType string // return type
}

func (s *Procedure) String() string {
	return ""
}

// ProcedureParam represents a stored procedure.
type ProcedureParam struct {
	Name *string // param name
	Type *string // param type
}

func (s *ProcedureParam) String() string {
	return ""
}

// View struct
type View struct {
	Table *Table
	Query string
}

func (s *View) String() string {
	return ""
}
