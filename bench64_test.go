package radixsort_test

import (
	"fmt"
	"runtime"
	"slices"
	"testing"

	rdxsort "github.com/loov/radixsort"
	"github.com/twotwotwo/sorts/sortutil"
)

func benchmarkStdLibSort64(b *testing.B, size int, mode string) {
	data := generateData[uint64](size, mode)
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append([]uint64{}, data...)
		slices.Sort(tmp)
	}
}

func BenchmarkStdLibSort64(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("StdLib64_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkStdLibSort64(b, size, mode)
			})
		}
	}
}

func benchmarkLoovRadixSort64(b *testing.B, size int, mode string) {
	data := generateData[uint64](size, mode)
	buf := make([]uint64, len(data))
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append([]uint64{}, data...)
		rdxsort.Uint64(tmp, buf)
	}
}

func BenchmarkLoovRadixSort64(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixSortByLoovUint64_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkLoovRadixSort64(b, size, mode)
			})
		}
	}
}

func benchmarkTwoTwoTwoRadixSort64(b *testing.B, size int, mode string) {
	data := generateData[uint64](size, mode)
	data = sortutil.Uint64Slice(data)
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append(sortutil.Uint64Slice{}, data...)
		tmp.Sort()
	}
}

func BenchmarkTwoTwoTwoRadixSort64(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixSortByTwoTwoTwoUint64_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkTwoTwoTwoRadixSort64(b, size, mode)
			})
		}
	}
}
