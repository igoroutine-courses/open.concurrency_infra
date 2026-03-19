package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var n int32

	// see race detector
	for range 1000 {
		go func() {
			atomic.AddInt32(&n, 1)
			n++
		}()
	}

	time.Sleep(time.Minute)
	fmt.Println(n)
}
