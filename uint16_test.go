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

func TestUint16(t *testing.T) {
	tests := []struct {
		name string
		in   []uint16
		want []uint16
	}{
		{
			name: "empty slice",
			in:   []uint16{},
			want: []uint16{},
		},
		{
			name: "single element",
			in:   []uint16{42},
			want: []uint16{42},
		},
		{
			name: "already sorted",
			in:   []uint16{1, 2, 3, 4, 5},
			want: []uint16{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []uint16{5, 4, 3, 2, 1},
			want: []uint16{1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []uint16{7, 3, 7, 1, 3, 1},
			want: []uint16{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []uint16{10, 200, 5, 150, 5, 42},
			want: []uint16{5, 5, 10, 42, 150, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint16, len(tt.in))
			data := append([]uint16{}, tt.in...)

			radixsort.Uint16(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Unit16(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestUint16MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]uint16, size)
		for i := range input {
			input[i] = uint16(rand.Int())
		}

		want := make([]uint16, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint16, len(input))
		data := append([]uint16(nil), input...)
		radixsort.Uint16(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkUint16(b *testing.B, size int, mode string) {
	data := generateData[uint16](size, mode)
	buf := make([]uint16, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]uint16{}, data...)
		radixsort.Uint16(tmp, buf)
	}
}

func BenchmarkUint16(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortUint16_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkUint16(b, size, mode)
			})
		}
	}
}
