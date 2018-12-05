package schema

import (
	"github.com/matthewmueller/errors"
	"github.com/matthewmueller/go-gen"
)

// Column struct
type Column struct {
	FieldOrdinal int     // field_ordinal
	Name         string  // column_name
	DataType     string  // data_type
	NotNull      bool    // not_null
	Comment      *string // description
	DefaultValue *string // default_value
	IsPrimaryKey bool    // is_primary_key
}

// Pascal case
func (c *Column) Pascal() string {
	return gen.Pascal(c.Name)
}

// Camel case
func (c *Column) Camel() string {
	return gen.Camel(c.Name)
}

// Snake case
func (c *Column) Snake() string {
	return gen.Snake(c.Name)
}

// Type of column
func (c *Column) Type(schema *Schema) (string, error) {
	dt, err := schema.Coerce.Type(c.DataType)
	if err != nil {
		return "", errors.Wrapf(err, "column %q", c.Name)
	}
	return dt, nil
}

// Nullable of column
func (c *Column) Nullable() bool {
	return !c.NotNull
}
