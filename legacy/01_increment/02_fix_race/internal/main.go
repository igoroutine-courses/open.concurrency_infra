package main

import (
	"fmt"
	"sync"
)

//go:noescape
func SyncAdd(addr *int64, delta int64)

func main() {
	var v int64
	wg := new(sync.WaitGroup)
	for range 1000 {
		wg.Go(func() {
			SyncAdd(&v, 1)
		})
	}

	wg.Wait()
	fmt.Println(v)
}
