package radixsort_test

import (
	"fmt"
	"runtime"
	"slices"
	"testing"

	rdxsort "github.com/loov/radixsort"
	"github.com/twotwotwo/sorts/sortutil"
)

func benchmarkStdLibSort32(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append([]uint32{}, data...)
		slices.Sort(tmp)
	}
}

func BenchmarkStdLibSort32(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("StdLib32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkStdLibSort32(b, size, mode)
			})
		}
	}
}

func benchmarkLoovRadixSort32(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	buf := make([]uint32, len(data))
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append([]uint32{}, data...)
		rdxsort.Uint32(tmp, buf)
	}
}

func BenchmarkLoovRadixSort32(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixSortByLoovUint32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkLoovRadixSort32(b, size, mode)
			})
		}
	}
}

func benchmarkTwoTwoTwoRadixSort32(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	data = sortutil.Uint32Slice(data)
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append(sortutil.Uint32Slice{}, data...)
		tmp.Sort()
	}
}

func BenchmarkTwoTwoTwoRadixSort32(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixSortByTwoTwoTwoUint32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkTwoTwoTwoRadixSort32(b, size, mode)
			})
		}
	}
}
