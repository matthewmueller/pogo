package template

import (
	gen "github.com/matthewmueller/go-gen"
	"github.com/matthewmueller/pogo/database"
)

var enum = gen.MustCompile("enum", `
	
`)

// Enum struct
type Enum struct {
	Package string
	Schema  *database.Schema
}

// Generate the base template
func (e *Enum) Generate() (string, error) {
	return enum(e)
}
