package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculate frequency concurrently
func ConcurrentFrequency(strings []string) FreqMap {
	c := make(chan FreqMap, 10)

	for _, s := range strings {
		go func(s string, c chan FreqMap) {
			c <- Frequency(s)
		}(s, c)
	}

	result := FreqMap{}
	for range strings {
		freqMap := <-c
		for k := range freqMap {
			result[k] += freqMap[k]
		}
	}

	return result
}
