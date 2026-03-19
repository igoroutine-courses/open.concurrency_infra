package main

import (
	"testing"
)

func BenchmarkMax(b *testing.B) {
	const (
		start      = 100_123
		iterations = 1000
	)

	b.Run("without delay", func(b *testing.B) {
		b.ReportAllocs()

		var value int64 = start
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := range iterations {
					StoreMax(&value, int64(i)+start)
				}
			}
		})
	})

	b.Run("with delay", func(b *testing.B) {
		b.ReportAllocs()

		var value int64 = start
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := range iterations {
					StoreMaxWithDelay(&value, int64(i)+start)
				}
			}
		})
	})
}
