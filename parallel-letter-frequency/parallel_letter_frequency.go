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

func ConcurrentFrequency(list []string) FreqMap {
	c := make(chan FreqMap, len(list))

	for _, s := range list {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	result := FreqMap{}

	for _ = range list {
		for r, f := range <-c {
			result[r] += f
		}
	}

	return result
}
