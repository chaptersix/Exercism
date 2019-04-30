package scrabble

import (
	"unicode"
)

var scoreCard = map[rune]int{
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// Score score the submitted word
func Score(sumbition string) int {
	total := 0
	for _, submitLetter := range sumbition {
		score, ok := scoreCard[unicode.ToUpper(submitLetter)]
		if ok {
			total += score
		} else {
			total++
		}
	}
	return total
}
