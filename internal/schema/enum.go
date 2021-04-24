package schema

import (
	"github.com/matthewmueller/gotext"
	"github.com/matthewmueller/text"
)

// Enum struct
type Enum struct {
	Name   string
	Values []*EnumValue
}

// Slug generates the slug case
func (e *Enum) Slug() string {
	return gotext.Lower(text.Slug(e.Name))
}

// Pascal generates the pascal case
func (e *Enum) Pascal() string {
	return gotext.Pascal(e.Name)
}

// Camel generates the camel case
func (e *Enum) Camel() string {
	return gotext.Camel(e.Name)
}

// EnumValue is an enum value
type EnumValue struct {
	Label string
	Order int
}

// Pascal case
func (v *EnumValue) Pascal() string {
	return gotext.Pascal(v.Label)
}
