package diffsquares

import "math"

// SquareOfSum square the sum of the first n natural numbers
func SquareOfSum(n int) int {

	squareOfSum, naturalNumber := 0.0, 1.0
	for i := 0; i < n; i++ {
		squareOfSum += naturalNumber
		naturalNumber += 1.0
	}

	return int(math.Pow(squareOfSum, 2))
}

// SumOfSquares sum the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sumOfSquares, naturalNumber := 0, 1.0

	for i := 0; i < n; i++ {
		sumOfSquares += int(math.Pow(naturalNumber, 2))
		naturalNumber += 1.0
	}

	return sumOfSquares
}

// Difference find the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
