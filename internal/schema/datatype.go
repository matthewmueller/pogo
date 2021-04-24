package schema

import "github.com/matthewmueller/gotext"

// DataType for a column
type DataType interface {
	String() string
}

// List type
type List struct{ DataType }

func (l *List) String() string {
	return "[]" + l.DataType.String()
}

// Null type
type Null struct{}

func (*Null) String() string {
	return "nil"
}

// String type
type String struct{}

func (*String) String() string {
	return "string"
}

// Boolean type
type Boolean struct{}

func (*Boolean) String() string {
	return "bool"
}

// Int type
type Int struct{}

func (*Int) String() string {
	return "int"
}

// Int64 type
type Int64 struct{}

func (*Int64) String() string {
	return "int64"
}

// Float32 type
type Float32 struct{}

func (*Float32) String() string {
	return "float32"
}

// Float64 type
type Float64 struct{}

func (*Float64) String() string {
	return "float64"
}

// DateTime type
type DateTime struct{}

func (*DateTime) String() string {
	return "time.Time"
}

// JSON type
type JSON struct{}

func (*JSON) String() string {
	return "json.RawMessage"
}

// Enumerable type
type Enumerable struct {
	Schema string
	Name   string
}

func (e *Enumerable) String() string {
	return "enum." + gotext.Pascal(e.Name)
}
