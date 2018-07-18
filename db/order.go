package db

import (
	"fmt"

	"github.com/matthewmueller/go-gen"
)

// OrderField struct
type OrderField struct {
	Name        string // column name
	Description string // column description
	DataType    string // column type (only for columns)
	FKReference string // fk name (only for fk references)
}

// Pascal case
func (o *OrderField) Pascal() string {
	return gen.Pascal(o.Name)
}

// Format the order by condition
func (o *OrderField) Format() string {
	return fmt.Sprintf(`"%s" %%s`, o.Name)
}
