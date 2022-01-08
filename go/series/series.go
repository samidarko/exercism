package series

// All returns a list of all substrings of s with length n
func All(n int, s string) (substrings []string) {
	sLen := len(s)
	offset := 0
	for {
		limit := offset + n
		if limit > sLen {
			break
		}
		substring := s[offset:limit]
		substrings = append(substrings, substring)
		offset++
	}

	return
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
