package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, errors.New("n is 0 or negative")
	}

	steps := 0

	if n == 1 {
		return steps, nil
	}

	for n != 1 {

		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		steps++

	}

	return steps, nil
}
