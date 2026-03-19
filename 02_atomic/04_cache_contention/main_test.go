package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"

	"golang.org/x/sys/cpu"
)

const cacheLineSize = unsafe.Sizeof(cpu.CacheLinePad{})

// false sharing
// [1],   [2],   [3]
// Core1           Core2  Core3

// [1.........],   [2],   [3]
// Core1           Core2  Core3

type data1 struct {
	value atomic.Int64
}

type data2 struct {
	value atomic.Int64
	_     [cacheLineSize - unsafe.Sizeof(atomic.Int64{})]byte // try to comment
}

func BenchmarkCacheContention(b *testing.B) {
	workers := runtime.GOMAXPROCS(-1)
	size := workers
	const iters = 1000

	b.Run("with contention", func(b *testing.B) {
		b.ReportAllocs()

		wg := new(sync.WaitGroup)
		s := make([]data1, size)
		b.ResetTimer()

		for b.Loop() {
			for w := range workers {
				wg.Go(func() {
					for range iters {
						s[w].value.Add(1)
					}
				})
			}

			wg.Wait()
		}
	})

	b.Run("without contention", func(b *testing.B) {
		b.ReportAllocs()

		wg := new(sync.WaitGroup)
		s := make([]data2, size)
		b.ResetTimer()

		for b.Loop() {
			for w := range workers {
				wg.Go(func() {
					for range iters {
						s[w].value.Add(1)
					}
				})
			}

			wg.Wait()
		}
	})
}
