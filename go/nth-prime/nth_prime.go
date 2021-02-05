package prime

// Nth given a number n, determine what the nth prime is.
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	var prime int

	for i := 2; n > 0; i++ {
		if isPrime(i) {
			prime = i
			n--
		}
	}

	return prime, true
}

func isPrime(n int) bool {
	m := n / 2
	for i := 2; i <= m; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
