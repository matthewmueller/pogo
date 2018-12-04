package schema

import gen "github.com/matthewmueller/go-gen"

// Enum struct
type Enum struct {
	Name   string
	Values []*EnumValue
}

// Slug generates the slug case
func (e *Enum) Slug() string {
	return gen.Lower(gen.Slug(e.Name))
}

// Pascal generates the pascal case
func (e *Enum) Pascal() string {
	return gen.Pascal(e.Name)
}

// Camel generates the camel case
func (e *Enum) Camel() string {
	return gen.Camel(e.Name)
}

// EnumValue is an enum value
type EnumValue struct {
	Label string
	Order int
}

// Pascal case
func (v *EnumValue) Pascal() string {
	return gen.Pascal(v.Label)
}
