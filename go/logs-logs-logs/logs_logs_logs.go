package logs

const (
	exclamationMark = '❗'
	magnifyingGlass = '🔍'
	sun             = '☀'
	recommendation  = "recommendation"
	search          = "search"
	weather         = "weather"
	defaultApp      = "default"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if r == exclamationMark {
			return recommendation
		}
		if r == magnifyingGlass {
			return search
		}
		if r == sun {
			return weather
		}
	}
	return defaultApp
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for i, r := range runes {
		if r == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
