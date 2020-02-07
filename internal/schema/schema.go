package schema

import (
	"fmt"
	"strconv"
	"strings"
)

// Schema struct
type Schema struct {
	Name   string
	Tables Tables
	Enums  Enums
	// Views      []*View
	// Procedures []*Procedure
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
	Foreigns Foreigns
	Uniques  Uniques
}

func (t *Table) String() string {
	var fields strings.Builder
	if len(t.Columns) > 0 {
		fields.WriteString("\n\t")
		fields.WriteString(t.Columns.String())
	}
	if t.Primary != nil {
		fields.WriteString(",\n\t")
		fields.WriteString(t.Primary.String())
	}
	if len(t.Foreigns) > 0 {
		fields.WriteString(",\n\t")
		fields.WriteString(t.Foreigns.String())
	}
	if len(t.Uniques) > 0 {
		fields.WriteString(",\n\t")
		fields.WriteString(t.Uniques.String())
	}
	fields.WriteString("\n")
	return fmt.Sprintf("create table %q.%q (%s);", t.Schema, t.Name, fields.String())
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
func (columns Columns) String() string {
	blocks := make([]string, len(columns))
	for i, column := range columns {
		blocks[i] = column.String()
	}
	return strings.Join(blocks, ",\n\t")
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
	return fmt.Sprintf(`create type %q.%q as enum (%s);`, e.Schema, e.Name, strings.Join(values, ", "))
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
	Name    string
	Columns Columns
}

func (p *Primary) String() string {
	if p == nil {
		return ""
	}
	columns := make([]string, len(p.Columns))
	for i, column := range p.Columns {
		columns[i] = strconv.Quote(column.Name)
	}
	return fmt.Sprintf(`constraint %q primary key (%s)`, p.Name, strings.Join(columns, ", "))
}

// Foreigns list
type Foreigns []*Foreign

func (foreigns Foreigns) String() string {
	blocks := make([]string, len(foreigns))
	for i, column := range foreigns {
		blocks[i] = column.String()
	}
	return strings.Join(blocks, ",\n\t")
}

// Foreign constraint
type Foreign struct {
	Name       string
	Columns    Columns
	RefSchema  string
	RefTable   string
	RefColumns Columns
}

func (f *Foreign) String() string {
	columns := make([]string, len(f.Columns))
	for i, column := range f.Columns {
		columns[i] = strconv.Quote(column.Name)
	}
	refColumns := make([]string, len(f.RefColumns))
	for i, column := range f.RefColumns {
		refColumns[i] = strconv.Quote(column.Name)
	}
	return fmt.Sprintf(`constraint %q foreign key (%s) references %q.%q (%s)`, f.Name, strings.Join(columns, ", "), f.RefSchema, f.RefTable, strings.Join(refColumns, ", "))
}

// Uniques list
type Uniques []*Unique

// Unique string
func (uniques Uniques) String() string {
	blocks := make([]string, len(uniques))
	for i, unique := range uniques {
		blocks[i] = unique.String()
	}
	return strings.Join(blocks, ",\n\t")
}

// Unique constraint
type Unique struct {
	Name    string
	Columns Columns
}

func (u *Unique) String() string {
	columns := make([]string, len(u.Columns))
	for i, column := range u.Columns {
		columns[i] = strconv.Quote(column.Name)
	}
	return fmt.Sprintf(`constraint %q unique (%s)`, u.Name, strings.Join(columns, ", "))
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
