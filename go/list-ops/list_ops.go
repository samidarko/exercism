package listops

type (
	IntList   []int
	binFunc   func(int, int) int
	predFunc  func(int) bool
	unaryFunc func(int) int
)

// Reverse an IntList
func (a IntList) Reverse() IntList {

	first := 0
	last := a.Length() - 1
	for first < last {
		a[first], a[last] = a[last], a[first]
		first++
		last--
	}

	return a
}

// Length gives IntList length
func (a IntList) Length() int {
	var length int

	for range a {
		length++
	}

	return length
}

// Append an IntList to another IntList
func (a IntList) Append(b IntList) IntList {
	aLength := a.Length()
	cLength := aLength + b.Length()
	c := make(IntList, cLength)
	for i, value := range a {
		c[i] = value
	}
	for i, value := range b {
		c[i+aLength] = value
	}
	return c
}

// Concat multiple IntList to an IntList
func (a IntList) Concat(bs []IntList) IntList {
	for _, b := range bs {
		a = a.Append(b)
	}
	return a
}

// Foldl fold left an IntList
func (a IntList) Foldl(fn binFunc, acc int) int {
	for _, value := range a {
		acc = fn(acc, value)
	}
	return acc
}

// Foldr fold right an IntList
func (a IntList) Foldr(fn binFunc, acc int) int {
	return a.Reverse().Foldl(func(b int, a int) int {
		return fn(a, b)
	}, acc)
}

// Map an IntList
func (a IntList) Map(fn unaryFunc) IntList {
	for i, value := range a {
		a[i] = fn(value)
	}
	return a
}

// Filter an IntList
func (a IntList) Filter(fn predFunc) IntList {
	c := make(IntList, 0)
	for _, value := range a {
		if fn(value) {
			c = c.Append(IntList{value})
		}
	}
	return c
}
