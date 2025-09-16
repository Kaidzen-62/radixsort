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

func TestInt64(t *testing.T) {
	tests := []struct {
		name string
		in   []int64
		want []int64
	}{
		{
			name: "empty slice",
			in:   []int64{},
			want: []int64{},
		},
		{
			name: "single element",
			in:   []int64{42},
			want: []int64{42},
		},
		{
			name: "single negative element",
			in:   []int64{-42},
			want: []int64{-42},
		},
		{
			name: "already sorted",
			in:   []int64{1, 2, 3, 4, 5},
			want: []int64{1, 2, 3, 4, 5},
		},
		{
			name: "already sorted negatives",
			in:   []int64{-5, -4, -3, -2, -1},
			want: []int64{-5, -4, -3, -2, -1},
		},
		{
			name: "already sorted mixed",
			in:   []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []int64{5, 4, 3, 2, 1},
			want: []int64{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order negatives",
			in:   []int64{-1, -2, -3, -4, -5},
			want: []int64{-5, -4, -3, -2, -1},
		},
		{
			name: "reverse order mixed",
			in:   []int64{-1, -2, -3, -4, -5, 5, 4, 3, 2, 1},
			want: []int64{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []int64{7, 3, 7, 1, 3, 1},
			want: []int64{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []int64{10, 126, 5, 100, 5, 42},
			want: []int64{5, 5, 10, 42, 100, 126},
		},
		{
			name: "random order mixed",
			in:   []int64{10, -5, -20, 5, 100, -5, 5, 2, 42, -126, 0},
			want: []int64{-126, -20, -5, -5, 0, 2, 5, 5, 10, 42, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint64, len(tt.in))
			data := append([]int64{}, tt.in...)

			radixsort.Int64(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Int64(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestInt64MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]int64, size)
		for i := range input {
			input[i] = int64(rand.Int())
		}

		want := make([]int64, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint64, len(input))
		data := append([]int64(nil), input...)
		radixsort.Int64(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkInt64(b *testing.B, size int, mode string) {
	data := generateData[int64](size, mode)
	buf := make([]uint64, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]int64{}, data...)
		radixsort.Int64(tmp, buf)
	}
}

func BenchmarkInt64(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortInt64_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkInt64(b, size, mode)
			})
		}
	}
}
