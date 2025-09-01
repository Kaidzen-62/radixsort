package archived

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"golang.org/x/exp/constraints"
)

var sizes = []int{100, 1000, 10_000, 100_000, 1_000_000}
var modes = []string{"random", "sorted", "reverse", "duplicates"}

func generateData[T constraints.Integer](n int, mode string) []T {
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

func benchmarkuint16b16(b *testing.B, size int, mode string) {
	data := generateData[uint16](size, mode)
	buf := make([]uint16, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]uint16{}, data...)
		radix16b16(tmp, buf)
	}
}

func BenchmarkRadix16b8(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortRadix16b16_ARCHIVE_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkuint16b16(b, size, mode)
			})
		}
	}
}

func benchmarkuint16b8Opt(b *testing.B, size int, mode string) {
	data := generateData[uint16](size, mode)
	buf := make([]uint16, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]uint16{}, data...)
		radix16b8Opt(tmp, buf)
	}
}

func BenchmarkRadix16b8Opt(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortRadix16b8Opt_ARCHIVE_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkuint16b8Opt(b, size, mode)
			})
		}
	}
}
