package summultiples

// SumMultiples find the sum of all the unique multiples of particular numbers up to but not including that number
func SumMultiples(limit int, divisors ...int) int {

	var sumMultiples int

	for n := 1; n < limit; n++ {
		if IsMultiple(n, divisors...) {
			sumMultiples += n
		}
	}

	return sumMultiples
}

// IsMultiple is true if n is multiple of one of the divisors
func IsMultiple(n int, divisors ...int) bool {
	for _, divisor := range divisors {
		if divisor != 0 && n%divisor == 0 {
			return true
		}
	}
	return false
}
