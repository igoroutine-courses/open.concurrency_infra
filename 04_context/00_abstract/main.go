package main

import (
	"context"
)

// Package context defines the Context type, which carries deadlines,
// cancellation signals, and other request-scoped values across API boundaries
// and between processes.
//
// Incoming requests to a server should create a [Context], and outgoing
// calls to servers should accept a Context. The chain of function
// calls between them must propagate the Context, optionally replacing
// it with a derived Context created using [WithCancel], [WithDeadline],
// [WithTimeout], or [WithValue].
//
// A Context may be canceled to indicate that work done on its behalf should stop.
// A Context with a deadline is canceled after the deadline passes.
// When a Context is canceled, all Contexts derived from it are also canceled.
//
// The [WithCancel], [WithDeadline], and [WithTimeout] functions take a
// Context (the parent) and return a derived Context (the child) and a
// [CancelFunc]. Calling the CancelFunc directly cancels the child and its
// children, removes the parent's reference to the child, and stops
// any associated timers. Failing to call the CancelFunc leaks the
// child and its children until the parent is canceled. The go vet tool
// checks that CancelFuncs are used on all control-flow paths.
//
// The [WithCancelCause], [WithDeadlineCause], and [WithTimeoutCause] functions
// return a [CancelCauseFunc], which takes an error and records it as
// the cancellation cause. Calling [Cause] on the canceled context
// or any of its children retrieves the cause. If no cause is specified,
// Cause(ctx) returns the same value as ctx.Err().
//
// Programs that use Contexts should follow these rules to keep interfaces
// consistent across packages and enable static analysis tools to check context
// propagation:
//
// Do not store Contexts inside a struct type; instead, pass a Context
// explicitly to each function that needs it.
//
//	func DoSomething(ctx context.Context, arg Arg) error {
//		// ... use ctx ...
//	}
//
// Do not pass a nil [Context], even if a function permits it. Pass [context.TODO]
// if you are unsure about which Context to use.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The same Context may be passed to functions running in different goroutines;
// Contexts are safe for simultaneous use by multiple goroutines.

//type Context interface {
//	// Возвращает время, когда операция будет оповещена о необходимости завершения
//	Deadline() (deadline time.Time, ok bool)
//
//	// Возвращает канал, который будет закрыт при необходимости завершить операцию
//	// Служит в качестве инструмента оповещения об отмене
//	Done() <-chan struct{}
//
//	// Если Done не закрыт - возвращает nil.
//	// Если Done закрыт, Err ошибку с объяснением причины:
//	// - Canceled - контекст был отменен
//	// - DeadlineExceeded - наступил дедлайн.
//	// После возврашения не nil ошибки Err всегда возвращает данную ошибку.
//	Err() error
//
//	// Позволяет получить произвольный объект из контекста
//	Value(key interface{}) interface{}
//}

func main() {
	context.Background()
}
