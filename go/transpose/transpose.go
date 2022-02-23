package transpose

import (
	"fmt"
	"strings"
)

// Transpose an input text to an output text
func Transpose(input []string) []string {
	result := make([]string, 0)

	for i, s := range input {
		for j, r := range s {
			if j < len(result) {
				result[j] = fmt.Sprint(result[j], strings.Repeat(" ", i-len(result[j])), string(r))
			} else {
				result = append(result, fmt.Sprint(strings.Repeat(" ", i), string(r)))
			}
		}
	}

	return result
}
