package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	runeToApp := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001F50D': "search",
		'\u2600':     "weather",
	}
	for _, v := range log {
		_, ok := runeToApp[v]
		if ok {
			return runeToApp[v]
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	for _, v := range log {
		if v == oldRune {
			log = strings.Replace(log, string(oldRune), string(newRune), -1)
		}
	}
	return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunesInLog := utf8.RuneCountInString(log)
	if numberOfRunesInLog > limit {
		return false
	}
	return true
}
