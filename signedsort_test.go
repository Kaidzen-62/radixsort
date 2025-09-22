package radixsort_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/Kaidzen-62/radixsort"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
)

func TestInt8(t *testing.T) {
	testSignedSort(t, radixsort.Int8, "Int8")
}

func TestInt8LargeRandom(t *testing.T) {
	testSignedSortLargeRandom(t, radixsort.Int8, "Int8")
}

func TestInt16(t *testing.T) {
	testSignedSort(t, radixsort.Int16, "Int16")
}

func TestInt16LargeRandom(t *testing.T) {
	testSignedSortLargeRandom(t, radixsort.Int16, "Int16")
}

func TestInt32(t *testing.T) {
	testSignedSort(t, radixsort.Int32, "Int32")
}

func TestInt32LargeRandom(t *testing.T) {
	testSignedSortLargeRandom(t, radixsort.Int32, "Int32")
}

func TestInt64(t *testing.T) {
	testSignedSort(t, radixsort.Int64, "Int64")
}

func TestInt64LargeRandom(t *testing.T) {
	testSignedSortLargeRandom(t, radixsort.Int64, "Int64")
}

func testSignedSort[T constraints.Signed, B constraints.Unsigned](t *testing.T, sortFunc func([]T, []B), sortFuncName string) {
	tests := []struct {
		name string
		in   []T
		want []T
	}{
		{
			name: "empty slice",
			in:   []T{},
			want: []T{},
		},
		{
			name: "single element",
			in:   []T{42},
			want: []T{42},
		},
		{
			name: "single negative element",
			in:   []T{-42},
			want: []T{-42},
		},
		{
			name: "already sorted",
			in:   []T{1, 2, 3, 4, 5},
			want: []T{1, 2, 3, 4, 5},
		},
		{
			name: "already sorted negatives",
			in:   []T{-5, -4, -3, -2, -1},
			want: []T{-5, -4, -3, -2, -1},
		},
		{
			name: "already sorted mixed",
			in:   []T{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: []T{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []T{5, 4, 3, 2, 1},
			want: []T{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order negatives",
			in:   []T{-1, -2, -3, -4, -5},
			want: []T{-5, -4, -3, -2, -1},
		},
		{
			name: "reverse order mixed",
			in:   []T{-1, -2, -3, -4, -5, 5, 4, 3, 2, 1},
			want: []T{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []T{7, 3, 7, 1, 3, 1},
			want: []T{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []T{10, 126, 5, 100, 5, 42},
			want: []T{5, 5, 10, 42, 100, 126},
		},
		{
			name: "random order mixed",
			in:   []T{10, -5, -20, 5, 100, -5, 5, 2, 42, -126, 0},
			want: []T{-126, -20, -5, -5, 0, 2, 5, 5, 10, 42, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]B, len(tt.in))
			data := append([]T{}, tt.in...)

			sortFunc(data, buf)

			got := data

			if !cmp.Equal(tt.want, got) {
				t.Errorf("case: %s; %s(%v) = %v, want %v", tt.name, sortFuncName, tt.in, got, tt.want)
			}
		})
	}
}

func testSignedSortLargeRandom[T constraints.Signed, B constraints.Unsigned](t *testing.T, sortFunc func([]T, []B), sortFuncName string) {
	size := 1_000_000
	input := make([]T, size)
	for i := range input {
		input[i] = T(rand.Int())
	}

	if slices.IsSorted(input) {
		t.Fatalf("tettible rand.rand")
	}

	data := append([]T(nil), input...)
	buf := make([]B, len(data))
	sortFunc(data, buf)

	if !slices.IsSorted(data) {
		t.Errorf("%s failed to sort data correctly", sortFuncName)
	}
}
