package radixsort

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

// func TestPositivesNegatives(t *testing.T) {
// 	var data []int64 = []int64{0, 1, 2, 3, 4, 5, -5, -4, -3, -2, -1}
// 	buf := make([]uint64, len(data))
//
// 	p := positives(data, buf)
// 	n := negatives(data, buf[len(p):])
//
// 	fmt.Printf("len: %d; cap: %d;\n", len(p), len(p))
// 	fmt.Printf("len: %d; cap: %d;\n", len(n), len(n))
//
// 	_ = buf
// }

func TestInt64ver2calls(t *testing.T) {
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
			in:   []int64{10, 200, 5, 150, 5, 42},
			want: []int64{5, 5, 10, 42, 150, 200},
		},
		{
			name: "random order mixed",
			in:   []int64{10, -5, -20, 5, 150, -5, 5, 2, 42, -200, 0},
			want: []int64{-200, -20, -5, -5, 0, 2, 5, 5, 10, 42, 150},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint64, len(tt.in))
			data := append([]int64{}, tt.in...)

			int64ver2calls(data, buf)

			if !reflect.DeepEqual(tt.want, data) {
				t.Errorf("int64ver2calls(%v) = %v, want %v", tt.in, data, tt.want)
			}
		})
	}
}

func TestInt64ver2callsMatchesStdlib(t *testing.T) {
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
		int64ver2calls(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}

// func TestInt64ver1call(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		in   []int64
// 		want []int64
// 	}{
// 		{
// 			name: "empty slice",
// 			in:   []int64{},
// 			want: []int64{},
// 		},
// 		{
// 			name: "single element",
// 			in:   []int64{42},
// 			want: []int64{42},
// 		},
// 		{
// 			name: "single negative element",
// 			in:   []int64{-42},
// 			want: []int64{-42},
// 		},
// 		{
// 			name: "already sorted",
// 			in:   []int64{1, 2, 3, 4, 5},
// 			want: []int64{1, 2, 3, 4, 5},
// 		},
// 		{
// 			name: "already sorted negatives",
// 			in:   []int64{-5, -4, -3, -2, -1},
// 			want: []int64{-5, -4, -3, -2, -1},
// 		},
// 		{
// 			name: "already sorted mixed",
// 			in:   []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
// 			want: []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
// 		},
// 		{
// 			name: "reverse order",
// 			in:   []int64{5, 4, 3, 2, 1},
// 			want: []int64{1, 2, 3, 4, 5},
// 		},
// 		{
// 			name: "reverse order negatives",
// 			in:   []int64{-1, -2, -3, -4, -5},
// 			want: []int64{-5, -4, -3, -2, -1},
// 		},
// 		{
// 			name: "reverse order mixed",
// 			in:   []int64{-1, -2, -3, -4, -5, 0, 5, 4, 3, 2, 1},
// 			want: []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
// 		},
// 		{
// 			name: "with duplicates",
// 			in:   []int64{7, 3, 7, 1, 3, 1},
// 			want: []int64{1, 1, 3, 3, 7, 7},
// 		},
// 		{
// 			name: "random order",
// 			in:   []int64{10, 200, 5, 150, 5, 42},
// 			want: []int64{5, 5, 10, 42, 150, 200},
// 		},
// 		{
// 			name: "random order mixed",
// 			in:   []int64{10, -5, -20, 5, 150, -5, 5, 2, 42, -200, 0},
// 			want: []int64{-200, -20, -5, -5, 0, 2, 5, 5, 10, 42, 150},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			buf := make([]uint64, len(tt.in))
// 			data := append([]int64{}, tt.in...)
//
// 			int64ver1call(data, buf)
//
// 			if !reflect.DeepEqual(tt.want, data) {
// 				t.Errorf("int64ver1call(%v) = %v, want %v", tt.in, data, tt.want)
// 			}
// 		})
// 	}
// }

func TestInt64ver1callMatchesStdlib(t *testing.T) {
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
		int64ver1call(data, buf)

		if !reflect.DeepEqual(want, data) {
			t.Fatalf("mismatch for size %d: data %v, want %v", size, data, want)
		}
	}
}
