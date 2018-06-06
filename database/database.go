package database

// Database interface
type Database interface {
	Introspect(schema string) (*Schema, error)
}

// Schema introspection
type Schema struct {
	Name   string
	Tables []*Table
	Enums  []*Enum
}

// Table represents table info.
type Table struct {
	Type        byte   // type
	Name        string // table_name
	ManualPk    bool   // manual_pk
	Columns     []*Column
	ForeignKeys []*ForeignKey
	Indexes     []*Index
}

// Column represents column info.
type Column struct {
	FieldOrdinal int     // field_ordinal
	Name         string  // column_name
	DataType     string  // data_type
	NotNull      bool    // not_null
	Comment      *string // description
	DefaultValue *string // default_value
	IsPrimaryKey bool    // is_primary_key

	// we add this on during introspection
	GoType string
}

// ForeignKey represents a foreign key.
type ForeignKey struct {
	ForeignKeyName string // foreign_key_name
	Name           string // column_name
	RefIndexName   string // ref_index_name
	RefTableName   string // ref_table_name
	RefColumnName  string // ref_column_name
	KeyID          int    // key_id
	SeqNo          int    // seq_no
	OnUpdate       string // on_update
	OnDelete       string // on_delete
	Match          string // match
}

// Index data
type Index struct {
	Name      string // index_name
	IsUnique  bool   // is_unique
	IsPrimary bool   // is_primary
	SeqNo     int    // seq_no
	Origin    string // origin
	IsPartial bool   // is_partial
	Columns   []*IndexColumn
}

// IndexColumn represents index column info.
type IndexColumn struct {
	SeqNo    int    // seq_no
	Cid      int    // cid
	Name     string // column_name
	DataType string

	// we add this on during introspection
	GoType string
}

// Enum struct
type Enum struct {
	Name   string
	Values []*Value
}

// Value represents a enum value.
type Value struct {
	Label string
	Order int
}

// PgColOrder represents index column order.
type PgColOrder struct {
	Ord string // ord
}
