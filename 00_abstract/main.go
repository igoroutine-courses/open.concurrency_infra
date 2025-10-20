package demo

import (
	"sync"
	"time"
)

type Map struct {
	once   *sync.Once
	mx     *sync.Mutex
	values map[any]any
}

func NewMap() *Map {
	return &Map{
		mx:   new(sync.Mutex),
		once: new(sync.Once),
	}
}

func (m *Map) Add(key, value any) {
	m.init()

	m.mx.Lock()
	m.values[key] = value
	m.mx.Unlock()
}

func (m *Map) init() {
	m.once.Do(func() {
		time.Sleep(time.Second)
		m.values = make(map[any]any)
	})
}
