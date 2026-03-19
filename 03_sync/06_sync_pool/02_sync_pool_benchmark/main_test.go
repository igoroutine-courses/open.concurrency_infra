package main

import (
	"sync"
	"testing"
)

type Person struct {
	name string
}

type PersonsPool struct {
	pool sync.Pool
}

func NewPersonsPool() *PersonsPool {
	return &PersonsPool{
		pool: sync.Pool{
			New: func() interface{} { return new(Person) },
		},
	}
}
func (p *PersonsPool) Get() *Person {
	return p.pool.Get().(*Person)
}

func (p *PersonsPool) Put(person *Person) {
	p.pool.Put(person)
}

var globalPerson *Person

func BenchmarkWithPool(b *testing.B) {
	pool := NewPersonsPool()
	for b.Loop() {
		person := pool.Get()
		person.name = "Igor"
		pool.Put(person)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	for b.Loop() {
		person := &Person{name: "Igor"}
		globalPerson = person
	}
}
