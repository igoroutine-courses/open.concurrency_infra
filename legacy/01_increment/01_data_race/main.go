package main

import (
	"fmt"
	"sync"
)

func main() {
	var v int64
	wg := new(sync.WaitGroup)
	for range 1000 {
		wg.Go(func() {
			v++
		})
	}

	wg.Wait()
	fmt.Println(v)
}
