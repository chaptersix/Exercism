package letter

import (
	"sync"
)

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

// ConcurrentFrequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func ConcurrentFrequency(wordList []string) FreqMap {
	m := &sync.Map{}
	mutex := &sync.WaitGroup{}
	mutex.Add(len(wordList))
	for _, word := range wordList {
		go func(w string) {
			for r := range w {
				value, loaded := m.LoadOrStore(r, 1)
				if loaded {
					v := value.(int)
					m.Store(r, v+1)
				}
			}
			mutex.Done()
		}(word)
	}
	mr := FreqMap{}
	mutex2 := &sync.Mutex{}
	mutex.Wait()
	m.Range(func(key, value interface{}) bool {

		k := key.(rune)
		v := value.(int)
		mutex2.Lock()
		mr[k] = v
		mutex2.Unlock()
		return true
	})
	return mr
}
