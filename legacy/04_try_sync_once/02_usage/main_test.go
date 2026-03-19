package demo

import (
	"math/rand/v2"
	"sync"
	"testing"
)

func TestSyncOnce(b *testing.T) {
	m := NewMap()
	wg := new(sync.WaitGroup)

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			m.Add(rand.N[int](10e9), "the nature of concurrency")
		}()
	}

	wg.Wait()
}
