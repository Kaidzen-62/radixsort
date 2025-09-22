package radixsort_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/Kaidzen-62/radixsort"
	"golang.org/x/exp/constraints"
)

func BenchmarkInt8(b *testing.B) {
	benchmarkSigned(b, radixsort.Int8, "Int8")
}

func BenchmarkInt16(b *testing.B) {
	benchmarkSigned(b, radixsort.Int16, "Int16")
}

func BenchmarkInt32(b *testing.B) {
	benchmarkSigned(b, radixsort.Int32, "Int32")
}

func BenchmarkInt64(b *testing.B) {
	benchmarkSigned(b, radixsort.Int64, "Int64")
}

func benchmarkSigned[T constraints.Signed, B constraints.Unsigned](b *testing.B, sortFunc func([]T, []B) error, sortFuncName string) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("Radixsort%s_%d_%s", sortFuncName, size, mode)
			}(), func(b *testing.B) {
				data := generateData[T](size, mode)
				buf := make([]B, len(data))
				runtime.GC()

				b.ResetTimer()
				for b.Loop() {
					tmp := append([]T{}, data...)
					err := sortFunc(tmp, buf)
					if err != nil {
						b.Fatalf("%s failed: %v", sortFuncName, err)
					}
				}
			})
		}
	}
}
