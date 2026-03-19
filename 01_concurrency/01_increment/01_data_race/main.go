package main

import (
	"fmt"
	"time"
)

func main() {
	var v int64

	for range 1000 {
		go func() {
			v++
		}()
	}

	// 945

	time.Sleep(time.Second * 5)
	fmt.Println(v)
}
