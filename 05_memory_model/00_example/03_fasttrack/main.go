package main

// The Go memory model specifies the conditions under which reads of a variable in one goroutine can be
// guaranteed to observe values produced by writes to the same variable in a different goroutine.

// If you must read the rest of this document to understand the behavior of your program, you are being too clever.

// Don't be clever.
