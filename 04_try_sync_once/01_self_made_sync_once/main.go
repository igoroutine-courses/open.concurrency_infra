package main

import (
	"sync/atomic"
)

type mySyncOnce struct {
	done atomic.Bool
}

func (m *mySyncOnce) Do(f func()) {
	if m.done.CompareAndSwap(false, true) {
		f()
	}
}
