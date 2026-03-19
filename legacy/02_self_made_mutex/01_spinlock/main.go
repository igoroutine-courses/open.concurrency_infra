package main

import (
	"runtime"
	"sync/atomic"
)

const (
	unlocked = iota
	locked
)

type spinLock struct {
	state atomic.Int64
}

func (s *spinLock) Lock() {
	for {
		if s.state.CompareAndSwap(unlocked, locked) {
			runtime.Gosched()
			return
		}
	}
}

func (s *spinLock) Unlock() {
	s.state.Store(unlocked)
}
