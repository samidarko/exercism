package allyourbase

import (
	"fmt"
)

// ConvertToBase convert to base
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	sum := 0
	for _, d := range inputDigits {
		sum += d
		if d < 0 || d >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
	}

	if sum == 0 {
		return []int{0}, nil
	}

	// convert to base 10
	tenBase := 0
	power := len(inputDigits) - 1
	for _, d := range inputDigits {
		tenBase += d * pow(inputBase, power)
		power--
	}

	// convert to output base
	result := make([]int, 0)
	for tenBase > 0 {
		remainder := tenBase % outputBase
		tenBase = tenBase / outputBase
		result = append([]int{remainder}, result...)
	}

	return result, nil
}

func pow(n, power int) int {
	if power == 0 {
		return 1
	}
	result := n
	for i := 2; i <= power; i++ {
		result *= n
	}
	return result
}
