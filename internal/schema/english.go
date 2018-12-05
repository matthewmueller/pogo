package schema

import (
	"strings"

	gen "github.com/matthewmueller/go-gen"
)

// Singular makes each segment singular
func singular(s string) string {
	segments := strings.Split(gen.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, gen.Pascal(gen.Singular(segment)))
	}
	return strings.Join(singles, "")
}

// Plural makes each segment singular
func plural(s string) string {
	segments := strings.Split(gen.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, gen.Pascal(gen.Plural(segment)))
	}
	return strings.Join(singles, "")
}
