package main

import "sync/atomic"

func StoreMax(
	addr1,
	addr2,
	addr3 *int64,
	candidate1,
	candidate2,
	candidate3 int64,
) {
	for {
		old1 := atomic.LoadInt64(addr1)
		old2 := atomic.LoadInt64(addr2)
		old3 := atomic.LoadInt64(addr3)

		if candidate1 <= old1 && candidate2 <= old2 && candidate3 <= old3 {
			return
		}

		if atomic.CompareAndSwapInt64(addr1, old1, candidate1) &&
			atomic.CompareAndSwapInt64(addr2, old2, candidate2) &&
			atomic.CompareAndSwapInt64(addr3, old3, candidate3) {
			return
		}
	}
}
