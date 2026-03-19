package main

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoreMax(t *testing.T) {
	for range 100_000 {
		wg := new(sync.WaitGroup)
		var v int64

		wg.Go(func() {
			StoreMax(&v, 10)
		})

		wg.Go(func() {
			StoreMax(&v, 15)
		})

		wg.Go(func() {
			StoreMax(&v, 20)
		})

		wg.Wait()
		require.EqualValues(t, 20, atomic.LoadInt64(&v))
	}
}
