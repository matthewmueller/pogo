package db

import (
	"github.com/matthewmueller/go-gen"
)

// ForeignKey struct
type ForeignKey struct {
	Name           string // column_name
	DataType       string // ref_data_type
	RefIndexName   string // ref_index_name
	RefTableName   string // ref_table_name
	RefColumnName  string // ref_column_name
	ForeignKeyName string // foreign_key_name
	KeyID          int    // key_id
	SeqNo          int    // seq_no
	OnUpdate       string // on_update
	OnDelete       string // on_delete
	Match          string // match
}

// Pascal case
func (f *ForeignKey) Pascal() string {
	return gen.Pascal(f.Name)
}

// Camel case
func (f *ForeignKey) Camel() string {
	return gen.Camel(f.Name)
}

// Snake case
func (f *ForeignKey) Snake() string {
	return gen.Snake(f.Name)
}

// Type of column
func (f *ForeignKey) Type(schema *Schema) (string, error) {
	return coerce(schema, f.DataType)
}
