package main

import (
	"math/rand/v2"
	"sync"
)

// We can use just global lock, Microsoft

type Client struct {
	mx     *sync.Mutex
	id     int64
	amount int64
}

func bankTransfer(left, right Client) {
	firstMx, secondMx := left.mx, right.mx

	if left.id > right.id {
		firstMx, secondMx = secondMx, firstMx
	}

	firstMx.Lock()
	defer firstMx.Unlock()

	secondMx.Lock()
	defer secondMx.Unlock()

	right.amount += 100
	left.amount -= 100
}

func main() {
	wg := sync.WaitGroup{}

	first := Client{
		mx:     new(sync.Mutex),
		id:     1,
		amount: 1000,
	}

	second := Client{
		mx:     new(sync.Mutex),
		id:     2,
		amount: 5000,
	}

	for range 5000 {
		wg.Go(func() {
			if rand.N[int](100)%2 == 0 {
				bankTransfer(first, second)
				return
			}

			bankTransfer(second, first)
		})
	}

	wg.Wait()
}
