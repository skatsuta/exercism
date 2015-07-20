package diffsquares

// SumOfSquares calculates the the sum of squares up to n.
func SumOfSquares(n int) int {
	// the formula for the sum of squares
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSums calculates the square of sums up to n.
func SquareOfSums(n int) int {
	// the formula for the sum of series
	sum := (1 + n) * n / 2
	return sum * sum
}

// Difference calculates the difference between the sum of squares and the square of sums.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
