package database

import gen "github.com/matthewmueller/go-gen"

// Enum struct
type Enum struct {
	Name   string
	Values []*EnumValue
}

// EnumValue is an enum value
type EnumValue struct {
	Label string
	Order int
}

// Slug generates the slug case
func (e *Enum) Slug() string {
	return gen.Lower(gen.Slug(e.Name))
}
