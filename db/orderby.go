package db

import (
	"fmt"

	"github.com/matthewmueller/go-gen"
)

// OrderByField struct
type OrderByField struct {
	Name        string // column name
	Description string // column description
	DataType    string // column type (only for columns)
	FKReference string // fk name (only for fk references)
}

// Pascal case
func (o *OrderByField) Pascal() string {
	return gen.Pascal(o.Name)
}

// Format the order by condition
func (o *OrderByField) Format() string {
	return fmt.Sprintf(`"%s" %%s`, o.Name)
}
