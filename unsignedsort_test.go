package radixsort_test

import (
	"errors"
	"math/rand"
	"slices"
	"testing"

	"github.com/Kaidzen-62/radixsort"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
)

func TestUint8(t *testing.T) {
	testUnsignedSort(t, radixsort.Uint8, "Uint8")
}

func TestUint16(t *testing.T) {
	testUnsignedSort(t, radixsort.Uint16, "Uint16")
}

func TestUint32(t *testing.T) {
	testUnsignedSort(t, radixsort.Uint32, "Uint32")
}

func TestUint64(t *testing.T) {
	testUnsignedSort(t, radixsort.Uint64, "Uint64")
}

func TestUint8LargeRandom(t *testing.T) {
	testUnsignedSortLargeRandom(t, radixsort.Uint8, "Uint8")
}

func TestUint16LargeRandom(t *testing.T) {
	testUnsignedSortLargeRandom(t, radixsort.Uint16, "Uint16")
}

func TestUint32LargeRandom(t *testing.T) {
	testUnsignedSortLargeRandom(t, radixsort.Uint32, "Uint32")
}

func TestUint64LargeRandom(t *testing.T) {
	testUnsignedSortLargeRandom(t, radixsort.Uint64, "Uint64")
}

func TestUint8BufferSize(t *testing.T) {
	testSortBufferSize(t, radixsort.Uint8, "Uint8")
}

func TestUint16BufferSize(t *testing.T) {
	testSortBufferSize(t, radixsort.Uint16, "Uint16")
}

func TestUint32BufferSize(t *testing.T) {
	testSortBufferSize(t, radixsort.Uint32, "Uint32")
}

func TestUint64BufferSize(t *testing.T) {
	testSortBufferSize(t, radixsort.Uint64, "Uint64")
}

func testUnsignedSort[T constraints.Unsigned](t *testing.T, sortFunc func([]T, []T) error, sortFuncName string) {
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
			name: "already sorted",
			in:   []T{1, 2, 3, 4, 5},
			want: []T{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []T{5, 4, 3, 2, 1},
			want: []T{1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []T{7, 3, 7, 1, 3, 1},
			want: []T{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []T{10, 200, 5, 150, 5, 42},
			want: []T{5, 5, 10, 42, 150, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]T, len(tt.in))
			data := append([]T{}, tt.in...)

			err := sortFunc(data, buf)
			if err != nil {
				t.Fatalf("%s failed: %v", sortFuncName, err)
			}

			got := data

			if !cmp.Equal(tt.want, got) {
				t.Errorf("case: %s; %s(%v) = %v, want %v", tt.name, sortFuncName, tt.in, got, tt.want)
			}
		})
	}
}

func testUnsignedSortLargeRandom[T constraints.Unsigned](t *testing.T, sortFunc func([]T, []T) error, sortFuncName string) {
	size := 1_000_000
	input := make([]T, size)
	for i := range input {
		input[i] = T(rand.Uint64())
	}

	if slices.IsSorted(input) {
		t.Fatalf("terrible rand.rand")
	}

	data := append([]T(nil), input...)
	buf := make([]T, len(data))

	err := sortFunc(data, buf)
	if err != nil {
		t.Fatalf("%s failed: %v", sortFuncName, err)
	}

	if !slices.IsSorted(data) {
		t.Errorf("%s failed to sort data correctly", sortFuncName)
	}
}

func testSortBufferSize[T constraints.Integer, B constraints.Unsigned](t *testing.T, sortFunc func([]T, []B) error, sortFuncName string) {
	const size = 10
	input := make([]T, size)
	for i := range input {
		input[i] = T(rand.Uint64())
	}

	tests := []struct {
		name       string
		bufSize    int
		wantErr    error
		shouldSort bool
	}{
		{
			name:       "BufferTooSmall",
			bufSize:    size - 5,
			wantErr:    radixsort.ErrInvalidBufferSize,
			shouldSort: false,
		},
		{
			name:       "BufferTooLarge",
			bufSize:    size + 5,
			wantErr:    nil,
			shouldSort: false,
		},
		{
			name:       "BufferCorrectSize",
			bufSize:    size,
			wantErr:    nil,
			shouldSort: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := append([]T(nil), input...)
			buf := make([]B, tc.bufSize)

			err := sortFunc(data, buf)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("%s: error = %v, want %v", sortFuncName, err, tc.wantErr)
			}

			if tc.shouldSort && !slices.IsSorted(data) {
				t.Errorf("%s: data is not sorted", sortFuncName)
			}
		})
	}
}
