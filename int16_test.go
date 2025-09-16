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

func TestInt16(t *testing.T) {
	tests := []struct {
		name string
		in   []int16
		want []int16
	}{
		{
			name: "empty slice",
			in:   []int16{},
			want: []int16{},
		},
		{
			name: "single element",
			in:   []int16{42},
			want: []int16{42},
		},
		{
			name: "single negative element",
			in:   []int16{-42},
			want: []int16{-42},
		},
		{
			name: "already sorted",
			in:   []int16{1, 2, 3, 4, 5},
			want: []int16{1, 2, 3, 4, 5},
		},
		{
			name: "already sorted negatives",
			in:   []int16{-5, -4, -3, -2, -1},
			want: []int16{-5, -4, -3, -2, -1},
		},
		{
			name: "already sorted mixed",
			in:   []int16{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: []int16{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []int16{5, 4, 3, 2, 1},
			want: []int16{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order negatives",
			in:   []int16{-1, -2, -3, -4, -5},
			want: []int16{-5, -4, -3, -2, -1},
		},
		{
			name: "reverse order mixed",
			in:   []int16{-1, -2, -3, -4, -5, 5, 4, 3, 2, 1},
			want: []int16{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []int16{7, 3, 7, 1, 3, 1},
			want: []int16{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []int16{10, 126, 5, 100, 5, 42},
			want: []int16{5, 5, 10, 42, 100, 126},
		},
		{
			name: "random order mixed",
			in:   []int16{10, -5, -20, 5, 100, -5, 5, 2, 42, -126, 0},
			want: []int16{-126, -20, -5, -5, 0, 2, 5, 5, 10, 42, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint16, len(tt.in))
			data := append([]int16{}, tt.in...)

			radixsort.Int16(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("Int16(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestInt16MatchesStdlib(t *testing.T) {
	for size := 1; size <= 1024; size *= 2 {
		input := make([]int16, size)
		for i := range input {
			input[i] = int16(rand.Int())
		}

		want := make([]int16, len(input))
		copy(want, input)
		slices.Sort(want)

		buf := make([]uint16, len(input))
		data := append([]int16(nil), input...)
		radixsort.Int16(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

func benchmarkInt16(b *testing.B, size int, mode string) {
	data := generateData[int16](size, mode)
	buf := make([]uint16, size)
	runtime.GC()

	b.ResetTimer()
	for b.Loop() {
		tmp := append([]int16{}, data...)
		radixsort.Int16(tmp, buf)
	}
}

func BenchmarkInt16(b *testing.B) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("RadixsortInt16_%d_%s", size, mode)
			}(), func(b *testing.B) {
				benchmarkInt16(b, size, mode)
			})
		}
	}
}
