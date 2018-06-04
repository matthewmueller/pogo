package template

import (
	gen "github.com/matthewmueller/go-gen"
	"github.com/matthewmueller/pogo/database"
)

var manyToMany = gen.MustCompile("many-to-many", `
	
`)

// ManyToMany struct
type ManyToMany struct {
	Schema *database.Schema
	Table  *database.Table
}

// Generate the base template
func (m *ManyToMany) Generate() (string, error) {
	return manyToMany(m)
}
