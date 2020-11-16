package anagram

import (
	"sort"
	"strings"
)

// Detect filter a list of candidates to return only the subject's anagrams
func Detect(subject string, candidates []string) []string {
	var anagrams []string
	for _, candidate := range candidates {
		if AreAnagrams(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

// AreAnagrams says true if left and right are anagrams
func AreAnagrams(left, right string) bool {
	left, right = strings.ToLower(left), strings.ToLower(right)
	if left == right {
		return false
	}
	l, r := []rune(left), []rune(right)
	sort.Slice(l, func(a, b int) bool { return l[a] < l[b] })
	sort.Slice(r, func(a, b int) bool { return r[a] < r[b] })
	return string(l) == string(r)
}
