package schema

import (
	"strings"

	text "github.com/matthewmueller/go-text"
)

// Singular makes each segment singular
func singular(s string) string {
	segments := strings.Split(text.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, text.Pascal(text.Singular(segment)))
	}
	return strings.Join(singles, "")
}

// Plural makes each segment singular
func plural(s string) string {
	segments := strings.Split(text.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, text.Pascal(text.Plural(segment)))
	}
	return strings.Join(singles, "")
}
