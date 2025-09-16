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

func TestInt32(t *testing.T) {
	tests := []struct {
		name string
		in   []int32
		want []int32
	}{
		{
			name: "empty slice",
			in:   []int32{},
			want: []int32{},
		},
		{
			name: "single element",
			in:   []int32{42},
			want: []int32{42},
		},
		{
			name: "single negative element",
			in:   []int32{-42},
			want: []int32{-42},
		},
		{
			name: "already sorted",
			in:   []int32{1, 2, 3, 4, 5},
			want: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "already sorted negatives",
			in:   []int32{-5, -4, -3, -2, -1},
			want: []int32{-5, -4, -3, -2, -1},
		},
		{
			name: "already sorted mixed",
			in:   []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []int32{5, 4, 3, 2, 1},
			want: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order negatives",
			in:   []int32{-1, -2, -3, -4, -5},
			want: []int32{-5, -4, -3, -2, -1},
		},
		{
			name: "reverse order mixed",
			in:   []int32{-1, -2, -3, -4, -5, 5, 4, 3, 2, 1},
			want: []int32{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []int32{7, 3, 7, 1, 3, 1},
			want: []int32{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []int32{10, 126, 5, 100, 5, 42},
			want: []int32{5, 5, 10, 42, 100, 126},
		},
		{
			name: "random order mixed",
			in:   []int32{10, -5, -20, 5, 100, -5, 5, 2, 42, -126, 0},
			want: []int32{-126, -20, -5, -5, 0, 2, 5, 5, 10, 42, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint32, len(tt.in))
			data := append([]int32{}, tt.in...)

			radixsort.Int32(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Int32(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestInt32MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]int32, size)
		for i := range input {
			input[i] = int32(rand.Int())
		}

		want := make([]int32, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint32, len(input))
		data := append([]int32(nil), input...)
		radixsort.Int32(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkInt32(b *testing.B, size int, mode string) {
	data := generateData[int32](size, mode)
	buf := make([]uint32, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]int32{}, data...)
		radixsort.Int32(tmp, buf)
	}
}

func BenchmarkInt32(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortInt32_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkInt32(b, size, mode)
			})
		}
	}
}
