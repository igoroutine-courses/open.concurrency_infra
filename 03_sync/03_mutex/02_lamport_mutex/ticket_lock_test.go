package mutexbench

import (
	"math/rand/v2"
	"sync"
	"testing"
)

func TestMutualExclusion(t *testing.T) {
	t.Parallel()

	v := make(map[int]int)
	mx := NewTicketLock()

	wg := new(sync.WaitGroup)
	for range 10 {
		wg.Go(func() {
			for range 1_000 {
				mx.Lock()
				v[rand.N[int](10e9)]++
				mx.Unlock()
			}
		})
	}

	wg.Wait()
}
