package summultiples

func SumMultiples(limit int, divisors ...int) int {

	var sumMultiples int

	for n := 1; n < limit; n++ {
		if IsMultiple(n, divisors...) {
			sumMultiples += n
		}
	}

	return sumMultiples
}

func IsMultiple(n int, divisors ...int) bool {
	for _, divisor := range divisors {
		if divisor != 0 && n%divisor == 0 {
			return true
		}
	}
	return false
}
