package diffsquares

// SquareOfSum square the sum of the first n natural numbers
func SquareOfSum(n int) int {

	squareOfSum, naturalNumber := 0, 1
	for i := 0; i < n; i++ {
		squareOfSum += naturalNumber
		naturalNumber += 1
	}

	return squareOfSum * squareOfSum
}

// SumOfSquares sum the squares of the first n natural numbers
func SumOfSquares(n int) int {
	sumOfSquares, naturalNumber := 0, 1

	for i := 0; i < n; i++ {
		sumOfSquares += naturalNumber * naturalNumber
		naturalNumber += 1
	}

	return sumOfSquares
}

// Difference find the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
