package schema

import (
	"github.com/matthewmueller/gotext"
)

// NewIndexColumn fn
func NewIndexColumn(
	name string,
	dataType DataType,
) *IndexColumn {
	return &IndexColumn{
		name,
		dataType,
	}
}

// IndexColumn represents index column info.
type IndexColumn struct {
	name     string
	dataType DataType
}

// Name of the index column
func (c *IndexColumn) Name() string {
	// for sqlite
	if c.name == "rowid" {
		return "RowID"
	}
	return c.name
}

// Pascal case
func (c *IndexColumn) Pascal() string {
	return gotext.Pascal(c.name)
}

// Camel case
func (c *IndexColumn) Camel() string {
	return gotext.Camel(c.name)
}

// Type of the column
func (c *IndexColumn) Type() string {
	return c.dataType.String()
}
