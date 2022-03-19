package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Set type here.
type Set map[string]bool

// New returns a new Set
func New() Set {
	return Set{}
}

// NewFromSlice returns a new Set from slice
func NewFromSlice(elems []string) Set {
	set := New()
	for _, elem := range elems {
		set.Add(elem)
	}
	return set
}

// String returns a string representation of Set
func (s Set) String() string {
	var output strings.Builder
	output.WriteString(`{`)
	if len(s) > 0 {
		output.WriteString(`"`)
	}
	elements := make([]string, 0)
	for element := range s {
		elements = append(elements, element)
	}
	output.WriteString(strings.Join(elements, `", "`))
	if len(s) > 0 {
		output.WriteString(`"`)
	}
	output.WriteString(`}`)
	return output.String()
}

// IsEmpty returns true if Set is empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if element belongs to Set
func (s Set) Has(elem string) bool {
	return s[elem]
}

// Add a new element to set
func (s Set) Add(elem string) {
	s[elem] = true
}

// Subset returns true if all elements of s1 are contained in s2
func Subset(s1, s2 Set) bool {
	for elem := range s1 {
		if s2.Has(elem) == false {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 do not have any elements in common
func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

// Equal returns true if s1 and s2 have the exact same elements
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

// Intersection returns a Set with elements contained in s1 and s2
func Intersection(s1, s2 Set) Set {
	set := New()
	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}
	for elem := range s1 {
		if s2.Has(elem) {
			set.Add(elem)
		}
	}
	return set
}

// Difference returns a Set with elements of s1 not contained in s2
func Difference(s1, s2 Set) Set {
	set := New()
	for elem := range s1 {
		if s2.Has(elem) == false {
			set.Add(elem)
		}
	}
	return set
}

// Union returns a Set with all the elements of s1 and s2
func Union(s1, s2 Set) Set {
	set := New()
	for elem := range s1 {
		set.Add(elem)
	}
	for elem := range s2 {
		set.Add(elem)
	}
	return set
}
