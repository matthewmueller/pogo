package testutil

import (
	"strings"
)

// Name of the test
func Name(t Test) string {
	name := t.Name
	if name == "" {
		name = Truncate(t.Func, 20)
		if i := strings.Index(t.Func, "("); i >= 0 {
			name = Truncate(t.Func, i)
		}
	}
	return name
}
