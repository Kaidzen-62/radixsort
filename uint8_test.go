package radixsort_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"slices"
	"testing"

	"github.com/Kaidzen-62/radixsort"
)

func TestUint8(t *testing.T) {
	tests := []struct {
		name string
		in   []uint8
		want []uint8
	}{
		{
			name: "empty slice",
			in:   []uint8{},
			want: []uint8{},
		},
		{
			name: "single element",
			in:   []uint8{42},
			want: []uint8{42},
		},
		{
			name: "already sorted",
			in:   []uint8{1, 2, 3, 4, 5},
			want: []uint8{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []uint8{5, 4, 3, 2, 1},
			want: []uint8{1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []uint8{7, 3, 7, 1, 3, 1},
			want: []uint8{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []uint8{10, 200, 5, 150, 5, 42},
			want: []uint8{5, 5, 10, 42, 150, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint8, len(tt.in))
			data := append([]uint8{}, tt.in...)

			radixsort.Uint8(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Unit8(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestUint8MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]uint8, size)
		for i := range input {
			input[i] = uint8(rand.Int())
		}

		want := make([]uint8, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint8, len(input))
		data := append([]uint8(nil), input...)
		radixsort.Uint8(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkUint8(b *testing.B, size int, mode string) {
	data := generateData[uint8](size, mode)
	buf := make([]uint8, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]uint8{}, data...)
		radixsort.Uint8(tmp, buf)
	}
}

func BenchmarkUint8(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortUint8_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkUint8(b, size, mode)
			})
		}
	}
}
