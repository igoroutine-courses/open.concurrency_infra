package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type CancelFunc = func()

var ErrContextCancelled = errors.New("context canceled")
var ErrTimeout = errors.New("timeout")

type MyContext struct {
	done chan struct{}
	err  error
	once sync.Once
	mu   sync.Mutex
}

// NewContext создает контекст с заданным таймаутом.
func NewContext(timeout time.Duration) (*MyContext, CancelFunc) {
	ctx := &MyContext{
		done: make(chan struct{}),
	}

	if timeout <= 0 {
		panic("invariant error")
	}

	go func() {
		select {
		case <-time.After(timeout):
			ctx.cancelWithError(ErrTimeout)
		case <-ctx.done:
		}
	}()

	return ctx, func() {
		ctx.cancelWithError(ErrContextCancelled)
	}
}

func (c *MyContext) Done() <-chan struct{} {
	return c.done
}

func (c *MyContext) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err
}

func (c *MyContext) Cancel() {
	c.cancelWithError(errors.New("context canceled"))
}

func (c *MyContext) cancelWithError(err error) {
	c.once.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.err = err
		close(c.done)
	})
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := run[int](ctx)
	fmt.Println(result, err)
}

var (
	ErrTimeoutExceeded = errors.New("timeout")
	ErrInvariantError  = errors.New("invariant error")
)

func run[T any](ctx context.Context) (T, error) {
	resCh := make(chan T)
	errCh := make(chan error, 1)

	go func() {
		defer close(resCh)
		defer close(errCh)

		select {
		case v, ok := <-slowFunction[T]():
			if !ok {
				errCh <- ErrInvariantError
				return
			}

			resCh <- v
		case <-ctx.Done():
			errCh <- ErrTimeoutExceeded
			return
		}
	}()

	return <-resCh, <-errCh // left-to-right order, so resCh is non-buffered
}

func slowFunction[T any]() <-chan T {
	ch := make(chan T, 1)

	go func() {
		defer close(ch)
		time.Sleep(time.Millisecond * rand.N[time.Duration](5000))

		ch <- *new(T)
	}()

	return ch
}
