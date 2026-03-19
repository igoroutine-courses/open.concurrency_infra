package main

import (
	"sync/atomic"
	"time"
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

	old1 := *addr1
	old2 := *addr2
	old3 := *addr3

	if candidate1 <= old1 && candidate2 <= old2 && candidate3 <= old3 {
		return
	}

	*addr1 = candidate1
	*addr2 = candidate2
	*addr3 = candidate3
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
		time.Sleep(time.Millisecond * 10)
	}
}

func (s *SpinLock) Unlock() {
	s.state.Store(unlocked)
}
