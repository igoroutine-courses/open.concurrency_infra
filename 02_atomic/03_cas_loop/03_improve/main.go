package main

import (
	"runtime"
	"sync/atomic"
)

// StoreMax атомарно обновляет *addr, записывая в неё candidate,
// если candidate > текущее значение. Необходимо использовать
// только функции пакета sync/atomic

// Пример: если в *addr сейчас 10, а candidate = 15,
// после вызова StoreMax(addr, 15) значение по адресу должно стать 15.
// Если candidate = 7, значение должно остаться 10.

func StoreMax(addr *int64, candidate int64) {
	for {
		old := atomic.LoadInt64(addr)

		if candidate <= old {
			return
		}

		if atomic.CompareAndSwapInt64(addr, old, candidate) {
			return
		}
	}
}

func StoreMaxWithDelay(addr *int64, candidate int64) {
	for {
		old := atomic.LoadInt64(addr)

		if candidate <= old {
			return
		}

		if atomic.CompareAndSwapInt64(addr, old, candidate) {
			return
		}

		runtime.Gosched() // or time.Sleep
	}
}
