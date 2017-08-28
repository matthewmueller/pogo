{{/*************************************************************************/}}
{{/* Variables */}}
{{/*************************************************************************/}}



{{/*************************************************************************/}}
{{/* Our Package */}}
{{/*************************************************************************/}}

package {{ .Settings.Package }}

{{/*************************************************************************/}}
{{/* Pogo marker */}}
{{/*************************************************************************/}}

// GENERATED BY POGO. DO NOT EDIT.

{{/*************************************************************************/}}
{{/* Database interface should work with database/sql.DB & database/sql.Tx */}}
{{/*************************************************************************/}}

// DB is the common interface for database operations that can be used with
// types from schema `{{ .Schema.Name }}`.
//
// This should work with database/sql.DB and database/sql.Tx.
type DB interface {
	Exec(string, ...interface{}) (pgx.CommandTag, error)
	Query(string, ...interface{}) (*pgx.Rows, error)
	QueryRow(string, ...interface{}) *pgx.Row
}

{{/*************************************************************************/}}
{{/* Public customizable logging interface */}}
{{/*************************************************************************/}}

// Log provides the log func used by generated queries.
var Log = func(string, ...interface{}) {}

{{/*************************************************************************/}}
{{/* Public options for upserts (TODO: make this an enum) */}}
{{/*************************************************************************/}}

const (
  // UpsertDoNothing Do nothing if there's a conflict
  UpsertDoNothing = "DO NOTHING"
  // UpsertDoUpdate Perform an update when there's a conflict
  UpsertDoUpdate = "DO UPDATE"
)

{{/*************************************************************************/}}
{{/* Public helper function to slice our fields into SQL friendly inputs */}}
{{/*************************************************************************/}}

// Slice converts our columns into something the sql driver can understand
func Slice(columns map[string]interface{}, offset int) (c []string, i []string, v []interface{}) {
	n := offset + 1
	for col, val := range columns {
		c = append(c, `"`+col+`"`)
		i = append(i, "$"+strconv.Itoa(n))
		v = append(v, val)
		n++
	}
	return c, i, v
}
