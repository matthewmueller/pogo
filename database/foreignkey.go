package database

// ForeignKey struct
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
