package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	appNames := map[string]string{
		"U+2757":  "recommendation",
		"U+1F50D": "search",
		"U+2600":  "weather",
	}
	for _, char := range log {
		appName := appNames[fmt.Sprintf("%U", char)]
		if appName != "" {
			return appName
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, char := range log {
		if char == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(char)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
