package logs

import (
	"strings"
	"unicode/utf8"
)

func compareRune(r rune) string {
	switch r {
	case '❗':
		return "recommendation"
	case '🔍':
		return "search"
	case '☀':
		return "weather"
	default:
		return "default"
	}
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		c := compareRune(c)
		if c != "default" {
			return c
		}
	}
	return "default"
}

// Replace replaces all occurances of old with new, returning the modified log
// to the caller.
func Replace(log string, old, new rune) string {

	return strings.ReplaceAll(log, string(old), string(new))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
