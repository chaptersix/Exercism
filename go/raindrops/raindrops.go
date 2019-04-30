package raindrops

import "strconv"

// Convert converts an integer into a string
func Convert(num int) string {
	returnStr := ""
	if num%3 == 0 {
		returnStr += "Pling"
	}
	if num%5 == 0 {
		returnStr += "Plang"
	}

	if num%7 == 0 {
		returnStr += "Plong"
	}

	if returnStr == "" {
		returnStr = strconv.Itoa(num)
	}
	return returnStr
}
