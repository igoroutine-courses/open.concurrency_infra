package main

import (
	"sync/atomic"
)

// StoreMax атомарно обновляет *addr, записывая в неё candidate,
// если candidate > текущее значение. Необходимо использовать
// только функции пакета sync/atomic

// Пример: если в *addr сейчас 10, а candidate = 15,
// после вызова StoreMax(addr, 15) значение по адресу должно стать 15.
// Если candidate = 7, значение должно остаться 10.

// 0

// G1
// StoreMax(addr, 10)

// G2
// StoreMax(addr, 15)

func StoreMax(addr *int64, candidate int64) {
	for {
		value := atomic.LoadInt64(addr)

		if value >= candidate {
			return
		}

		if atomic.CompareAndSwapInt64(addr, value, candidate) {
			break
		}
	}
}

func StoreMaxInvalid(addr *int64, candidate int64) {
	value := atomic.LoadInt64(addr)

	if value >= candidate {
		return
	}

	atomic.StoreInt64(addr, candidate)
}
