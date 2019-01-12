package pogo

// Postgres Enum type
// https://www.postgresql.org/docs/11/datatype-enum.html
//
// Enumerated (enum) types are data types that comprise a static,
// ordered set of values. They are equivalent to the enum types
// supported in a number of programming languages. An example of an
// enum type might be the days of the week, or a set of status values
// for a piece of data.
//
// Enum labels are case sensitive, so 'happy' is not the same as 'HAPPY'.
// White space in the labels is significant too.
//
// An enum value occupies four bytes on disk. The length of an enum value's
// textual label is limited by the NAMEDATALEN setting compiled into PostgreSQL;
// in standard builds this means at most 63 bytes.
//
// The translations from internal enum values to textual labels are kept in
// the system catalog pg_enum. Querying this catalog directly can be useful.

// Enum type
type Enum struct {
	Values []*EnumValue
}

// EnumValue type
type EnumValue struct {
}
