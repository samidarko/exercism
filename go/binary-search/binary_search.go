package binarysearch

// SearchInts binary search algorithm
func SearchInts(list []int, key int) int {
	for left, right := 0, len(list)-1; left <= right; {
		index := (left + right) / 2
		switch {
		case list[index] < key:
			left = index + 1
		case list[index] > key:
			right = index - 1
		default:
			return index
		}
	}
	return -1
}
