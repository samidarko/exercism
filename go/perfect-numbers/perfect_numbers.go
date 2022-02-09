package perfect

import (
	"errors"
)

// Classification type
type Classification int

func GetAliquotSum(n int64) int64 {
	var aliquotSum int64
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			aliquotSum += i
		}
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

	aliquotSum := GetAliquotSum(n)
	switch {
	case aliquotSum > n:
		return ClassificationAbundant, nil
	case aliquotSum < n || aliquotSum == 1:
		return ClassificationDeficient, nil
	default:
		return ClassificationPerfect, nil
	}
}
