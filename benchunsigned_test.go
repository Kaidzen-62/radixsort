package radixsort_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/Kaidzen-62/radixsort"
	"golang.org/x/exp/constraints"
)

func BenchmarkUint8(b *testing.B) {
	benchmarkUnsigned(b, radixsort.Uint8, "Uint8")
}

func BenchmarkUint16(b *testing.B) {
	benchmarkUnsigned(b, radixsort.Uint16, "Uint16")
}

func BenchmarkUint32(b *testing.B) {
	benchmarkUnsigned(b, radixsort.Uint32, "Uint32")
}

func BenchmarkUint64(b *testing.B) {
	benchmarkUnsigned(b, radixsort.Uint64, "Uint64")
}

func benchmarkUnsigned[T constraints.Unsigned](b *testing.B, sortFunc func([]T, []T), sortFuncName string) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("Radixsort%s_%d_%s", sortFuncName, size, mode)
			}(), func(b *testing.B) {
				data := generateData[T](size, mode)
				buf := make([]T, len(data))
				runtime.GC()

				b.ResetTimer()
				for b.Loop() {
					tmp := append([]T{}, data...)
					sortFunc(tmp, buf)
				}
			})
		}
	}
}
