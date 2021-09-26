package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	m := strings.Fields(line)
	return strings.Join(m[1:], " ")
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return utf8.RuneCountInString(Message(line))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	m := strings.Fields(line)
	r := strings.Replace(m[0], "[", "", 1)
	r = strings.Replace(r, "]", "", 1)
	r = strings.Replace(r, ":", "", 1)
	return strings.ToLower(r)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	m := Message(line)
	ll := LogLevel(line)
	return fmt.Sprintf("%s (%s)", m, ll)
}
