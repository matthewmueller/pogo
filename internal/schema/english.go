package schema

import (
	"strings"

	"github.com/matthewmueller/text"
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

// Plural makes each segment plural
func plural(s string) string {
	segments := strings.Split(text.Snake(s), "_")
	var singles []string
	for _, segment := range segments {
		singles = append(singles, text.Pascal(text.Plural(segment)))
	}
	return strings.Join(singles, "")
}
