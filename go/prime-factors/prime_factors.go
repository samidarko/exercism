package prime

// Factors compute the prime factors of a given natural number
func Factors(input int64) []int64 {

	factors := make([]int64, 0)

	for factor := int64(2); input > 1; {
		if input%factor == 0 {
			factors = append(factors, factor)
			input /= factor
			continue
		}
		factor++
	}
	return factors
}
