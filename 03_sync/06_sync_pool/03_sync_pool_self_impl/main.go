package main

import "sync"

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type linkedList[T any] struct {
	head *Node[T]
}

func newLinkedList[T any]() *linkedList[T] {
	return new(linkedList[T])
}

func (l *linkedList[T]) push(node *Node[T]) {
	node.next = l.head
	l.head = node
}

func (l *linkedList[T]) pop() *Node[T] {
	if l.head != nil {
		node := l.head
		l.head = l.head.next
		return node
	}

	return nil
}

type Pool[T any] struct {
	ctr func() T

	mtx  sync.Mutex
	list *linkedList[T]
}

func NewPool[T any](ctr func() T) *Pool[T] {
	return &Pool[T]{
		ctr:  ctr,
		list: newLinkedList[T](),
	}
}

func (l *Pool[T]) Get() *Node[T] {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	node := l.list.pop()
	if node == nil {
		node = &Node[T]{
			Value: l.ctr(),
		}
	}

	return node
}

func (l *Pool[T]) Put(node *Node[T]) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.list.push(node)
}
