package main

import "sync"

type Repository[K comparable, V any] struct {
	m sync.Map
}

func (c *Repository[K, V]) Load(key K) (V, bool) {
	value, found := c.m.Load(key)
	return value.(V), found
}

func (c *Repository[K, V]) Store(key K, value V) {
	c.m.Store(key, value)
}
