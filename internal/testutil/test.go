package testutil

// Test struct
type Test struct {
	Skip   bool
	Schema string
	Name   string
	Before string
	After  string
	Func   string
	Expect string
	Error  string
}
