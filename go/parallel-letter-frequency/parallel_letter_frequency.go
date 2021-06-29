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
	c := make(chan FreqMap, len(strings))

	for _, s := range strings {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	result := FreqMap{}
	for i := 0; i < cap(c); i++ {
		freqMap := <-c
		for k := range freqMap {
			result[k] = result[k] + freqMap[k]
		}
	}

	return result
}
