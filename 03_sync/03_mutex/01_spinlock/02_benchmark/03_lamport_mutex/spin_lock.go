package mutexbench

import (
	"sync"
	"sync/atomic"
)

var _ sync.Locker = (*spinLock)(nil)

func NewSpinLock() *spinLock {
	return &spinLock{}
}

const (
	unlocked = iota
	locked
)

type spinLock struct {
	state atomic.Int64
}

func (s *spinLock) Lock() {
	for !s.state.CompareAndSwap(unlocked, locked) {
	}
}

func (s *spinLock) Unlock() {
	s.state.Store(unlocked)
}
