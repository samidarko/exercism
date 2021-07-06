package armstrong

import (
	"math"
	"strconv"
)

// IsNumber returns true if n is an Armstrong number
func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	power := float64(len(s))
	var result float64
	for _, r := range s {
		result += math.Pow(float64(r-'0'), power)
	}
	return int(result) == n
}
