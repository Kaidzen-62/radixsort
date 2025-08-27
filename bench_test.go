package radixsort_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"slices"
	"testing"

	rdxsort "github.com/loov/radixsort"
	"github.com/twotwotwo/sorts/sortutil"
	"golang.org/x/exp/constraints"
)

var sizes = []int{100, 1000, 10_000, 100_000, 1_000_000}
var modes = []string{"random", "sorted", "reverse", "duplicates"}

func generateData[T 
                  .Integer](n int, mode string) []T {
	data := make([]T, n)
	switch mode {
	case "sorted":
		for i := range data {
			data[i] = T(i)
		}
	case "reverse":
		for i := range data {
			data[i] = T(n - i)
		}
	case "duplicates":
		for i := range data {
			data[i] = T(i % 66)
		}
	default: // "random"
		for i := range data {
			data[i] = T(rand.Intn(n * 10))
		}
	}
	return data
}

func benchmarkStdLibSort(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append([]uint32{}, data...)
		slices.Sort(tmp)
	}
}

func BenchmarkStdLibSort(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("StdLib_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkStdLibSort(b, size, mode)
			})
		}
	}
}

func benchmarkLoovRadixSort(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	buf := make([]uint32, len(data))
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append([]uint32{}, data...)
		rdxsort.Uint32(tmp, buf)
	}
}

func BenchmarkLoovRadixSort(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixSortByLoovUint32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkLoovRadixSort(b, size, mode)
			})
		}
	}
}

func benchmarkTwoTwoTwoRadixSort(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	data = sortutil.Uint32Slice(data)
	runtime.GC()
	b.ResetTimer()

	for b.Loop() {
		tmp := append(sortutil.Uint32Slice{}, data...)
		tmp.Sort()
	}
}

func BenchmarkTwoTwoTwoRadixSort(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixSortByTwoTwoTwoUint32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkTwoTwoTwoRadixSort(b, size, mode)
			})
		}
	}
}