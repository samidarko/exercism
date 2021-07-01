package sieve

// Number is struct that allows prime numbers marking
type Number struct {
	value   int
	isPrime bool
}

// Sieve find all the primes from 2 up to a given number
func Sieve(limit int) []int {

	// create `numbers` an int series from 2 to limit (hence -1)
	numbers := make([]Number, limit-1, limit-1)
	for i := range numbers {
		numbers[i] = Number{isPrime: true, value: i + 2}
	}

	// apply algorithm
	for _, number := range numbers {
		if number.isPrime {
			for n := number.value + number.value; n <= len(numbers)+1; n = n + number.value {
				numbers[n-2].isPrime = false
			}
		}
	}

	// compute result, a list of primes if found
	result := make([]int, 0)
	for _, number := range numbers {
		if number.isPrime {
			result = append(result, number.value)
		}
	}

	return result
}
