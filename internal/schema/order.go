package schema

import (
	"fmt"

	"github.com/matthewmueller/gotext"
)

func newOrder(
	name string,
	comment *string,
	dataType DataType,
) *OrderField {
	return &OrderField{
		name,
		comment,
		dataType,
	}
}

// OrderField struct
type OrderField struct {
	name     string   // column name
	comment  *string  // column comment
	dataType DataType // column type (only for columns)
}

// Name fn
func (o *OrderField) Name() string {
	return fmt.Sprintf("%q", o.name)
}

// Pascal case
func (o *OrderField) Pascal() string {
	return gotext.Pascal(o.name)
}

// Format the order by condition
func (o *OrderField) Format() string {
	return fmt.Sprintf(`"%s" %%s`, o.name)
}

// Type of order field
func (o *OrderField) Type() string {
	return o.dataType.String()
}
