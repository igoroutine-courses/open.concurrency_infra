package main

import (
	"fmt"
	"sync"
)

// (1, 0): 989248 times
// (0, 1): 10481 times
// (0, 0): 270 times
// (1, 1): 1 times

func main() {
	var (
		wg      sync.WaitGroup
		results = make(map[[2]int]int)
	)

	const iterations = 1_000_000
	for i := 0; i < iterations; i++ {
		var x, y int32
		var r1, r2 int32

		// G1
		wg.Go(func() {
			x = 1
			r1 = y
		})

		// G2
		wg.Go(func() {
			y = 1
			r2 = x
		})

		wg.Wait()
		results[[2]int{int(r1), int(r2)}]++
	}

	fmt.Println("Observed outcomes:")
	for k, v := range results {
		fmt.Printf("(%d, %d): %d times\n", k[0], k[1], v)
	}
}
