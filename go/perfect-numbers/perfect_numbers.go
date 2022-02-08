package perfect

import (
	"errors"
)

// Define the Classification type here.

type Classification int

func GetDivisors(n int) []int {
	divisors := []int{1}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func GetAliquotSum(divisors []int) int64 {
	var aliquotSum int64
	for _, d := range divisors {
		aliquotSum += int64(d)
	}
	return aliquotSum
}

const (
	ClassificationAbundant Classification = iota
	ClassificationDeficient
	ClassificationPerfect
)

var ErrOnlyPositive = errors.New("not a positive")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}

	divisors := GetDivisors(int(n))
	aliquotSum := GetAliquotSum(divisors)
	if aliquotSum > n {
		return ClassificationAbundant, nil
	}
	if aliquotSum < n || aliquotSum == 1 {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
