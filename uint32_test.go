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

func TestUint32(t *testing.T) {
	tests := []struct {
		name string
		in   []uint32
		want []uint32
	}{
		{
			name: "empty slice",
			in:   []uint32{},
			want: []uint32{},
		},
		{
			name: "single element",
			in:   []uint32{42},
			want: []uint32{42},
		},
		{
			name: "already sorted",
			in:   []uint32{1, 2, 3, 4, 5},
			want: []uint32{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []uint32{5, 4, 3, 2, 1},
			want: []uint32{1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []uint32{7, 3, 7, 1, 3, 1},
			want: []uint32{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []uint32{10, 200, 5, 150, 5, 42},
			want: []uint32{5, 5, 10, 42, 150, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint32, len(tt.in))
			data := append([]uint32{}, tt.in...)

			radixsort.Uint32(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Unit32(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestUint32MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]uint32, size)
		for i := range input {
			input[i] = uint32(rand.Int())
		}

		want := make([]uint32, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint32, len(input))
		data := append([]uint32(nil), input...)
		radixsort.Uint32(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkUint32(b *testing.B, size int, mode string) {
	data := generateData[uint32](size, mode)
	buf := make([]uint32, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]uint32{}, data...)
		radixsort.Uint32(tmp, buf)
	}
}

func BenchmarkUint32(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortUint32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkUint32(b, size, mode)
			})
		}
	}
}
