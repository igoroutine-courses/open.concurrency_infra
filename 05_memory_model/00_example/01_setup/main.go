package main

import "fmt"

var value string
var done bool

func setup() {
	value = "Hello, @igoroutine!"
	done = true
}

func main() {
	go setup()

	for !done {
	}

	fmt.Println(value)
}
