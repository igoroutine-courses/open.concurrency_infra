package demo

import (
	"sync"
	"sync/atomic"
	"time"
)

type mySyncOnce struct {
	done atomic.Bool
}

func (m *mySyncOnce) Do(f func()) {
	if m.done.CompareAndSwap(false, true) {
		f()
	}
}

type Map struct {
	once   *mySyncOnce
	mx     *sync.Mutex
	values map[interface{}]interface{}
}

func NewMap() *Map {
	return &Map{
		mx:   new(sync.Mutex),
		once: new(mySyncOnce),
	}
}

func (m *Map) Add(key, value interface{}) {
	m.init()

	m.mx.Lock()
	m.values[key] = value
	m.mx.Unlock()
}

func (m *Map) init() {
	m.once.Do(func() {
		time.Sleep(time.Second)
		m.values = make(map[interface{}]interface{})
	})
}
