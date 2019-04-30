package prime

import "math"

// Nth returns the nth prime of num
func Nth(nth int) (num int, valid bool) {
	if nth < 1 {
		return 0, false
	}
	primeCount := 0
	currentPrime := 0
	for i := 2; primeCount < nth; i++ {
		if isPrime(i) {
			primeCount++
			currentPrime = i
		}
	}
	return currentPrime, true
}

func isPrime(num int) bool {
	if num%2 == 0 {
		if num != 2 {
			return false
		}
		return true
	}
	if num%3 == 0 {
		if num != 3 {
			return false
		}
		return true
	}
	limit := int(math.Sqrt(float64(num))) + 1
	for i := 3; i < limit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
