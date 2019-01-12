package pogo

// Database struct
type Database struct {
	Schemas []*Schema
}

// TestString returns a string used for testing purposes
func (d *Database) TestString() string {
	return ""
}
