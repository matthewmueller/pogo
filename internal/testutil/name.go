package testutil

import (
	"strings"
)

// Name of the test
func Name(t Test) string {
	name := t.Name
	if name == "" {
		name = Truncate(t.Func, 80)
	}
	name = strings.ReplaceAll(name, " ", "")
	return name
}
