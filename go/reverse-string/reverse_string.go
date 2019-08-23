package reverse

import "strings"

//Reverse gives the inverse of the the provided string
func Reverse(forward string) string {
	var sb strings.Builder
	runes := []rune(forward)
	sb.Grow(len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}
