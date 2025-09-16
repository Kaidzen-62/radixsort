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

func TestInt8(t *testing.T) {
	tests := []struct {
		name string
		in   []int8
		want []int8
	}{
		{
			name: "empty slice",
			in:   []int8{},
			want: []int8{},
		},
		{
			name: "single element",
			in:   []int8{42},
			want: []int8{42},
		},
		{
			name: "single negative element",
			in:   []int8{-42},
			want: []int8{-42},
		},
		{
			name: "already sorted",
			in:   []int8{1, 2, 3, 4, 5},
			want: []int8{1, 2, 3, 4, 5},
		},
		{
			name: "already sorted negatives",
			in:   []int8{-5, -4, -3, -2, -1},
			want: []int8{-5, -4, -3, -2, -1},
		},
		{
			name: "already sorted mixed",
			in:   []int8{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: []int8{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []int8{5, 4, 3, 2, 1},
			want: []int8{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order negatives",
			in:   []int8{-1, -2, -3, -4, -5},
			want: []int8{-5, -4, -3, -2, -1},
		},
		{
			name: "reverse order mixed",
			in:   []int8{-1, -2, -3, -4, -5, 5, 4, 3, 2, 1},
			want: []int8{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []int8{7, 3, 7, 1, 3, 1},
			want: []int8{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []int8{10, 126, 5, 100, 5, 42},
			want: []int8{5, 5, 10, 42, 100, 126},
		},
		{
			name: "random order mixed",
			in:   []int8{10, -5, -20, 5, 100, -5, 5, 2, 42, -126, 0},
			want: []int8{-126, -20, -5, -5, 0, 2, 5, 5, 10, 42, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint8, len(tt.in))
			data := append([]int8{}, tt.in...)

			radixsort.Int8(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Int8(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestInt8MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]int8, size)
		for i := range input {
			input[i] = int8(rand.Int())
		}

		want := make([]int8, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint8, len(input))
		data := append([]int8(nil), input...)
		radixsort.Int8(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkInt8(b *testing.B, size int, mode string) {
	data := generateData[int8](size, mode)
	buf := make([]uint8, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]int8{}, data...)
		radixsort.Int8(tmp, buf)
	}
}

func BenchmarkInt8(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortInt8_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkInt8(b, size, mode)
			})
		}
	}
}
