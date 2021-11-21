package flatten

// Flatten Take a nested list and return a single flattened list with all values except nil/null
func Flatten(input interface{}) []interface{} {

	result := make([]interface{}, 0)

	switch element := input.(type) {
	case []interface{}:
		for _, e := range element {
			result = append(result, Flatten(e)...)
		}
	case interface{}:
		result = append(result, element)
	}

	return result
}
