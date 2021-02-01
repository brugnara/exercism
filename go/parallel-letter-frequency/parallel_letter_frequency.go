package letter

import "sync"

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

// ConcurrentFrequency computes concurrently the frequency
func ConcurrentFrequency(xs []string) FreqMap {
	m := FreqMap{}
	var wg sync.WaitGroup
	chn := make(chan FreqMap)
	done := make(chan bool)

	for _, s := range xs {
		wg.Add(1)

		go func(s string) {
			defer wg.Done()
			m := FreqMap{}
			for _, r := range s {
				m[r]++
			}
			chn <- m
		}(s)
	}

	go func() {
		for c := range chn {
			for k, v := range c {
				m[k] += v
			}
		}
		done <- true
	}()

	wg.Wait()
	close(chn)

	<-done

	return m
}
