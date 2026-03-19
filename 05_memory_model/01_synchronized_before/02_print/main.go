package main

import (
	"fmt"
	"time"
)

var value string

func doPrint() {
	fmt.Println(value)
}

func main() {
	value = "Hello, @igoroutine!"
	go doPrint()

	time.Sleep(time.Minute)
}

// The go statement that starts a new goroutine is synchronized before the start of the goroutine's execution.
