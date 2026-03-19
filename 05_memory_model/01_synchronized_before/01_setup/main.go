package main

import (
	"fmt"
	"sync/atomic"
)

var value string
var done atomic.Bool

func setup() {
	value = "Hello, @igoroutine!"
	done.Store(true)
}

func main() {
	go setup()

	for !done.Load() {
	}

	fmt.Println(value)
}

// The APIs in the sync/atomic package are collectively “atomic operations” that can be used to synchronize
// the execution of different goroutines. If the effect of an atomic operation A is observed by atomic operation B,
// then A is synchronized before B

// - Если A sequenced before B в одной горутине, то A happens before B
// - Если A synchronized before B (через механизмы синхронизации), то A happens before B
// - Если A happens before B, а B happens before C, то A happens before C (транзитивность)
