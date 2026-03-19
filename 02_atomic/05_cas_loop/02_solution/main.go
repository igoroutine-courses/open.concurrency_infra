package main

import (
	"runtime"
	"sync/atomic"
)

var lock = new(SpinLock)

func StoreMax(
	addr1,
	addr2,
	addr3 *int64,
	candidate1,
	candidate2,
	candidate3 int64,
) {
	lock.Lock()
	defer lock.Unlock()

	old1 := atomic.LoadInt64(addr1)
	old2 := atomic.LoadInt64(addr2)
	old3 := atomic.LoadInt64(addr3)

	if candidate1 <= old1 && candidate2 <= old2 && candidate3 <= old3 {
		return
	}

	atomic.StoreInt64(addr1, candidate1)
	atomic.StoreInt64(addr2, candidate2)
	atomic.StoreInt64(addr3, candidate3)
}

const (
	unlocked = iota
	locked
)

type SpinLock struct {
	state atomic.Int64
}

func (s *SpinLock) Lock() {
	for {
		if s.state.CompareAndSwap(unlocked, locked) {
			return
		}

		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	s.state.Store(unlocked)
}
