package acmp_concurrent

import (
	"sync"
	"test/acmp"
)

func Difficulties(urls []string) map[string]float64 {
	difficulties := make(map[string]float64)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result := acmp.Difficulty(url)
			mu.Lock()
			difficulties[url] = result
			mu.Unlock()
		}(url)
	}
	wg.Wait()
	return difficulties
}
