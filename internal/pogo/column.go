package pogo

// Column struct
type Column struct {
	Name         string   // column_name
	DataType     DataType // data_type
	NotNull      bool     // not_null
	Comment      *string  // description
	DefaultValue *string  // default_value
}
