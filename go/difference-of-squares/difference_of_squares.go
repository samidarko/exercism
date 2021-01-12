package diffsquares

// SquareOfSum square the sum of the first n natural numbers
func SquareOfSum(n int) int {

	squareOfSum := (n * (n + 1)) / 2

	return squareOfSum * squareOfSum
}

// SumOfSquares sum the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference find the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
