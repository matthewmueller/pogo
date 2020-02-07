package tests

// Test struct
type Test struct {
	Title   string
	Up      string
	Down    string
	QueryGo string
	QueryTS string

	// expects
	Schema string
	Expect string
	Error  string
}
