package isogram

import "strings"

// IsIsogram test if the provided string is an Isogram
func IsIsogram(query string) bool {
	m := make(map[rune]bool)
	for _, letter := range strings.ToUpper(query) {
		if letter == ' ' || letter == '-' {
			continue
		}
		if m[letter] {
			return false
		}
		m[letter] = true
	}
	return true
}
