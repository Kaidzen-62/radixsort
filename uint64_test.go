package radixsort_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"slices"
	"testing"

	"radixsort"
)

func TestUint64(t *testing.T) {
	tests := []struct {
		name string
		in   []uint64
		want []uint64
	}{
		{
			name: "empty slice",
			in:   []uint64{},
			want: []uint64{},
		},
		{
			name: "single element",
			in:   []uint64{42},
			want: []uint64{42},
		},
		{
			name: "already sorted",
			in:   []uint64{1, 2, 3, 4, 5},
			want: []uint64{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []uint64{5, 4, 3, 2, 1},
			want: []uint64{1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []uint64{7, 3, 7, 1, 3, 1},
			want: []uint64{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []uint64{10, 200, 5, 150, 5, 42},
			want: []uint64{5, 5, 10, 42, 150, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint64, len(tt.in))
			data := append([]uint64{}, tt.in...)

			radixsort.Uint64(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Unit64(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestUint64MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]uint64, size)
		for i := range input {
			input[i] = uint64(rand.Int())
		}

		want := make([]uint64, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint64, len(input))
		data := append([]uint64(nil), input...)
		radixsort.Uint64(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkUint64(b *testing.B, size int, mode string) {
	data := generateData[uint64](size, mode)
	buf := make([]uint64, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]uint64{}, data...)
		radixsort.Uint64(tmp, buf)
	}
}

func BenchmarkUint64(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortUint64_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkUint64(b, size, mode)
			})
		}
	}
}
