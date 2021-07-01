package sieve

type Number struct {
	value   int
	isPrime bool
}

func Sieve(limit int) []int {

	numbers := make([]Number, limit-1, limit-1)

	for i, _ := range numbers {
		numbers[i] = Number{isPrime: true, value: i + 2}
	}

	for _, number := range numbers {
		if number.isPrime {
			for n := number.value + number.value; n <= len(numbers)+1; n = n + number.value {
				numbers[n-2].isPrime = false
			}
		}
	}

	result := make([]int, 0)

	for _, number := range numbers {
		if number.isPrime {
			result = append(result, number.value)
		}
	}

	return result
}
