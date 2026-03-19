package main

import (
	"fmt"
	"sync"
)

type Data struct {
	First  int
	Second int
}

func main() {
	var data Data
	values := make([]int, 2)

	wg := new(sync.WaitGroup)

	// [G1 G1 G1 G1]...[G2, G2, G2, G2]

	wg.Go(func() {
		data.First = 27
		values[0] = 10
	})

	wg.Go(func() {
		data.Second = 28
		values[1] = 11
	})

	wg.Wait()

	fmt.Println(data.First)
	fmt.Println(data.Second)

	fmt.Println(values[0])
	fmt.Println(values[1])
	// Will we see it here?
}
