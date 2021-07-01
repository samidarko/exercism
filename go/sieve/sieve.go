package sieve

// Number is struct that allows prime numbers marking
type Number struct {
	value   int
	isPrime bool
}

// Sieve find all the primes from 2 up to a given number
func Sieve(limit int) []int {

	// create a sequence of numbers from 2 to limit (hence -1)
	sequence := make([]Number, limit-1, limit-1)
	for i := range sequence {
		sequence[i] = Number{isPrime: true, value: i + 2}
	}

	// apply algorithm
	primes := make([]int, 0)
	for _, number := range sequence {
		if number.isPrime {
			primes = append(primes, number.value)
			for n := number.value + number.value; n <= len(sequence)+1; n = n + number.value {
				sequence[n-2].isPrime = false
			}
		}
	}

	return primes
}
