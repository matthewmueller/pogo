package schema

import gen "github.com/matthewmueller/go-gen"

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

// Integer type
type Integer struct{}

func (*Integer) String() string {
	return "int"
}

// Float type
type Float struct{}

func (*Float) String() string {
	return "float"
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
	return "enum." + gen.Pascal(e.Name)
}
