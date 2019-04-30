package summultiples

// SumMultiples returns the sum of the multiples of num
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for num := range divisors {
			if num >= 1 && num%i == 0 {
				sum += i
			}
		}
	}
	return sum
}
