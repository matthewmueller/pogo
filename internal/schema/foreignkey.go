package schema

import "github.com/matthewmueller/gotext"

// NewForeignKey fn
func NewForeignKey(
	name string,
	dataType DataType,
) *ForeignKey {
	return &ForeignKey{
		name,
		dataType,
	}
}

// ForeignKey struct
type ForeignKey struct {
	name     string // column_name
	dataType DataType
}

// Pascal case
func (f *ForeignKey) Pascal() string {
	// for sqlite
	if f.name == "rowid" {
		return "RowID"
	}
	return gotext.Pascal(f.name)
}

// Camel case
func (f *ForeignKey) Camel() string {
	return gotext.Camel(f.name)
}

// Snake case
func (f *ForeignKey) Snake() string {
	return gotext.Snake(f.name)
}

// Type of column
func (f *ForeignKey) Type() string {
	return f.dataType.String()
}
