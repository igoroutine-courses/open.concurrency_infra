package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// C++ - std::mutex
// Java - ReentrantLock / synchronized
// C# - lock / Monitor / Mutex
// Python - threading.Lock
// Rust- std::sync::Mutex
// Kotlin - ReentrantLock / kotlinx.coroutines.Mutex
// JS/TS - нет стандартного, только библиотеки

// G0 Lock()
// G0 Unlock()
// G1 Lock()
// G1 Unlock()

type Data struct {
	Field1 string
	Field2 int
	Field3 int
}

func main() {
	d := Data{}

	wg := new(sync.WaitGroup)
	mx := &SpinLock{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		mx.Lock()
		fmt.Println("G0 Lock()")

		defer func() {
			fmt.Println("G0 Unlock()")
			mx.Unlock()
		}()

		d.Field1 = "concurrency"
		d.Field2 = 42
		d.Field3 = 1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mx.Lock()
		fmt.Println("G1 Lock()")

		defer func() {
			fmt.Println("G1 Unlock()")
			mx.Unlock()
		}()

		d.Field1 = "concurrency"
		d.Field2 = 42
		d.Field3 = 1
	}()

	wg.Wait()
}

const (
	unlocked = iota
	locked
)

type SpinLock struct {
	state atomic.Int64
}

func (s *SpinLock) Lock() {
	for !s.state.CompareAndSwap(unlocked, locked) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	s.state.Store(unlocked)
}
