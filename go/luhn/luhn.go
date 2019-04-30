package luhn

import "strings"

// Valid checks for a valid card number
func Valid(cardNumber string) bool {
	cardNumber = strings.Replace(cardNumber, " ", "", -1)
	cardNumberLength := len(cardNumber)
	if cardNumberLength <= 1 {
		return (false)
	}
	sum := 0
	double := false
	for i := cardNumberLength - 1; i >= 0; i-- {
		num := int(cardNumber[i] - '0')
		if num < 0 || num > 9 {
			return false
		}
		if double {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		double = !double
		sum += num
	}
	return sum%10 == 0
}
