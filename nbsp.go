package nbsp

import (
	"regexp"
)

// Repl represents the string which replaces the spaces in the
var Repl = "\u00A0"

// Setup contains data of the expressions to find and replace
type Setup struct {
	Regex       string
	Replacement string
}

// Replace will retrieve the original string and return the modified one
func Replace(s string, list []Setup) string {
	for {
		// replace all cases by its replacement string
		gen := s
		for i := 0; i < len(list); i++ {
			gen = regexp.MustCompile(list[i].Regex).ReplaceAllString(gen, list[i].Replacement)
		}

		// do the comparison
		if s == gen {
			return s
		}

		// assign new value
		s = gen
	}
}
