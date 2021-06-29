package prime

import (
	"math/big"
)

// Factors compute the prime factors of a given natural number
func Factors(input int64) []int64 {

	factors := make([]int64, 0)

	for i := int64(2); input > 1; {
		if big.NewInt(i).ProbablyPrime(0) && input%i == 0 {
			factors = append(factors, i)
			input /= i
			continue
		}
		i++
	}
	return factors
}
