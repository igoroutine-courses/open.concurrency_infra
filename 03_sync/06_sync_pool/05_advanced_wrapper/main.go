package mypool

import (
	"sync"
)

type MyPool[T any] struct {
	pool *sync.Pool
}

func NewMyPool[T any](newFunc func() T) *MyPool[T] {
	return &MyPool[T]{
		pool: &sync.Pool{
			New: func() any {
				return newFunc()
			},
		},
	}
}

func (p *MyPool[T]) Get() T {
	return p.pool.Get().(T)
}

func (p *MyPool[T]) Put(x T) {
	p.pool.Put(x)
}
