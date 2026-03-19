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
