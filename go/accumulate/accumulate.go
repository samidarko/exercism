// Package accumulate accumulate provide accumulation operation
package accumulate

// Accumulate apply operation to each item of list
func Accumulate(list []string, operation func(string) string) []string {

	for i, val := range list {
		list[i] = operation(val)
	}

	return list
}
