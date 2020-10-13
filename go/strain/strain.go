// Package strain is helpful
package strain

// Ints is an alias for []int
type Ints []int

// Strings is an alias for []string
type Strings []string

// Lists is an alias for []Ints
type Lists []Ints

// Keep return kept values fulfilling predicate
func (xs Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, x := range xs {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// Keep return kept values fulfilling predicate
func (xs Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, x := range xs {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// Keep return kept values fulfilling predicate
func (xs Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, x := range xs {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// Discard return discarded values unfulfilling predicate
func (xs Ints) Discard(predicate func(int) bool) Ints {
	return xs.Keep(func(n int) bool {
		return !predicate(n)
	})
}

// Discard return discarded values unfulfilling predicate
func (xs Strings) Discard(predicate func(string) bool) Strings {
	return xs.Keep(func(n string) bool {
		return !predicate(n)
	})
}

// Discard return discarded values unfulfilling predicate
func (xs Lists) Discard(predicate func([]int) bool) Lists {
	return xs.Keep(func(n []int) bool {
		return !predicate(n)
	})
}
