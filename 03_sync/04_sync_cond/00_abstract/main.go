package main

// Cond implements a condition variable, a rendezvous point
// for goroutines waiting for or announcing the occurrence
// of an event.
//
// Each Cond has an associated Locker L (often a [*Mutex] or [*RWMutex]),
// which must be held when changing the condition and
// when calling the [Cond.Wait] method.
//
// A Cond must not be copied after first use.
//
// In the terminology of [the Go memory model], Cond arranges that
// a call to [Cond.Broadcast] or [Cond.Signal] “synchronizes before” any Wait call
// that it unblocks.
//
// For many simple use cases, users will be better off using channels than a
// Cond (Broadcast corresponds to closing a channel, and Signal corresponds to
// sending on a channel).
