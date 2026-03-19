package main

import (
	"fmt"
	"time"
)

// LDADDAL
// https://developer.arm.com/documentation/111108/2025-09/Base-Instructions/LDADD--LDADDA--LDADDAL--LDADDL--Atomic-add-on-word-or-doubleword-?lang=en

//go:noescape
func SyncAdd(addr *int64, delta int64)

func main() {
	var v int64

	for range 1000 {
		go func() {
			SyncAdd(&v, 1) // v++
		}()
	}

	// 1000

	time.Sleep(time.Second * 5)
	fmt.Println(v)
}
