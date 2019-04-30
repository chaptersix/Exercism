package diffsquares

// SquareOfSum computes the square of  all natural  numbers up to and including the input
func SquareOfSum(n int) int {
	// summation  from 1 to n of i
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares computes the sum of the  squares of all natural numbers up to and including the input
func SumOfSquares(n int) int {
	// summation from  1 to n of i^2
	return int((n * (n + 1) * ((2 * n) + 1)) / 6)
}

// Difference computes the difference between the Square of  the Sum of the input minus the Sum of the Sum of the Squares of the input
func Difference(a int) int {
	return SquareOfSum(a) - SumOfSquares(a)
}
