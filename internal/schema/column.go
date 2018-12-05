package schema

import (
	"github.com/matthewmueller/go-gen"
)

// NewColumn fn
func NewColumn(
	name string,
	alias string,
	dataType DataType,
	notNull bool,
	comment *string,
	defaultValue *string,
	isPrimaryKey bool,
) *Column {
	return &Column{
		name,
		alias,
		dataType,
		notNull,
		comment,
		defaultValue,
		isPrimaryKey,
	}
}

// Column struct
type Column struct {
	name         string // column_name
	alias        string
	dataType     DataType // data_type
	notNull      bool     // not_null
	comment      *string  // description
	defaultValue *string  // default_value
	isPrimaryKey bool     // is_primary_key
}

// Name of the column
func (c *Column) Name() string {
	return c.name
}

// Pascal case
func (c *Column) Pascal() string {
	if c.alias != "" {
		return gen.Pascal(c.alias)
	}
	return gen.Pascal(c.name)
}

// Camel case
func (c *Column) Camel() string {
	if c.alias != "" {
		return gen.Camel(c.alias)
	}
	return gen.Camel(c.name)
}

// Snake case
func (c *Column) Snake() string {
	return gen.Snake(c.name)
}

// Type of column
func (c *Column) Type() string {
	return c.dataType.String()
}

// Nullable of column
func (c *Column) Nullable() bool {
	return !c.notNull
}
