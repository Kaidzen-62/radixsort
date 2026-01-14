package radixsort_test

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"github.com/Kaidzen-62/radixsort"
	"github.com/google/go-cmp/cmp"
)

func TestGenericInt8(t *testing.T) {
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
			in:   []int8{10, math.MaxInt8, 5, 100, 5, 42},
			want: []int8{5, 5, 10, 42, 100, math.MaxInt8},
		},
		{
			name: "random order mixed",
			in:   []int8{10, -5, -20, 5, math.MaxInt8, -5, 5, 2, 42, math.MinInt8, 0},
			want: []int8{math.MinInt8, -20, -5, -5, 0, 2, 5, 5, 10, 42, math.MaxInt8},
		},
	}

	key := func(a int8) int8 {
		return a
	}

	sortFuncName := "Generic[int8](data, buf []int8)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]int8, len(tt.in))
			data := append([]int8{}, tt.in...)

			err := radixsort.Generic(data, buf, key)
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

func TestGenericInt64(t *testing.T) {
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
			in:   []int64{10, math.MaxInt64, 5, 100, 5, 42},
			want: []int64{5, 5, 10, 42, 100, math.MaxInt64},
		},
		{
			name: "random order mixed",
			in:   []int64{10, -5, -20, 5, math.MaxInt64, -5, 5, 2, 42, math.MinInt64, 0},
			want: []int64{math.MinInt64, -20, -5, -5, 0, 2, 5, 5, 10, 42, math.MaxInt64},
		},
	}

	key := func(a int64) int64 {
		return a
	}

	sortFuncName := "Generic[int64](data, buf []int64)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]int64, len(tt.in))
			data := append([]int64{}, tt.in...)

			err := radixsort.Generic(data, buf, key)
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

func TestGenericUint64(t *testing.T) {
	tests := []struct {
		name string
		in   []uint64
		want []uint64
	}{
		{
			name: "empty slice",
			in:   []uint64{},
			want: []uint64{},
		},
		{
			name: "single element",
			in:   []uint64{42},
			want: []uint64{42},
		},
		{
			name: "already sorted",
			in:   []uint64{1, 2, 3, 4, 5},
			want: []uint64{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order",
			in:   []uint64{5, 4, 3, 2, 1},
			want: []uint64{1, 2, 3, 4, 5},
		},
		{
			name: "with duplicates",
			in:   []uint64{7, 3, 7, 1, 3, 1},
			want: []uint64{1, 1, 3, 3, 7, 7},
		},
		{
			name: "random order",
			in:   []uint64{10, math.MaxUint64, 5, 100, 5, 42},
			want: []uint64{5, 5, 10, 42, 100, math.MaxUint64},
		},
	}

	key := func(a uint64) uint64 {
		return a
	}

	sortFuncName := "Generic[uint64](data, buf []uint64)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]uint64, len(tt.in))
			data := append([]uint64{}, tt.in...)

			err := radixsort.Generic(data, buf, key)
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

func TestGenericUint64Random(t *testing.T) {
	size := 1_000_000
	input := make([]uint64, size)
	for i := range input {
		input[i] = uint64(rand.Uint64())
	}

	if slices.IsSorted(input) {
		t.Fatalf("terrible rand.rand")
	}

	data := append([]uint64(nil), input...)
	buf := make([]uint64, len(data))
	key := func(a uint64) uint64 {
		return a
	}

	err := radixsort.Generic(data, buf, key)
	if err != nil {
		t.Fatalf("%s failed: %v", "Generic[uint64](randomData, buf, key)", err)
	}

	if !slices.IsSorted(data) {
		t.Errorf("%s failed to sort data correctly", "Generic[uint64](randomData, buf, key)")
	}
}

func TestGenericIntRandom(t *testing.T) {
	size := 4_000_000
	input := make([]int, size)
	for i := range input {
		input[i] = rand.Int()
	}

	if slices.IsSorted(input) {
		t.Fatalf("terrible rand.rand")
	}

	data := append([]int(nil), input...)
	buf := make([]int, len(data))
	key := func(a int) int {
		return a
	}

	err := radixsort.Generic(data, buf, key)
	if err != nil {
		t.Fatalf("%s failed: %v", "Generic[int](randomData, buf, key)", err)
	}

	if !slices.IsSorted(data) {
		t.Errorf("%s failed to sort data correctly", "Generic[int](randomData, buf, key)")
	}
}

func TestGenreicFloat32(t *testing.T) {
	tests := []struct {
		name string
		in   []float32
		want []float32
	}{
		{
			name: "empty slice",
			in:   []float32{},
			want: []float32{},
		},
		{
			name: "single element",
			in:   []float32{4.2},
			want: []float32{4.2},
		},
		{
			name: "single negative element",
			in:   []float32{-4.2},
			want: []float32{-4.2},
		},
		{
			name: "already sorted",
			in:   []float32{1.1, 2.02, 3.003, 4.4, 5.05},
			want: []float32{1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "already sorted negatives",
			in:   []float32{-5.05, -4.4, -3.003, -2.02, -1.1},
			want: []float32{-5.05, -4.4, -3.003, -2.02, -1.1},
		},
		{
			name: "already sorted mixed",
			in:   []float32{-5.05, -4.4, -3.003, -2.02, -1.1, 0, 1.1, 2.02, 3.003, 4.4, 5.05},
			want: []float32{-5.05, -4.4, -3.003, -2.02, -1.1, 0, 1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "reverse order",
			in:   []float32{5.05, 4.4, 3.003, 2.02, 1.1},
			want: []float32{1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "reverse order negatives",
			in:   []float32{-1.1, -2.02, -3.003, -4.4, -5.05},
			want: []float32{-5.05, -4.4, -3.003, -2.02, -1.1},
		},
		{
			name: "reverse order mixed",
			in:   []float32{-1.1, -2.02, -3.003, -4.4, -5.05, 5.05, 4.4, 3.003, 2.02, 1.1},
			want: []float32{-5.05, -4.4, -3.003, -2.02, -1.1, 1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "with duplicates",
			in:   []float32{7.007, 3.003, 7.007, 1.1, 3.003, 1.1},
			want: []float32{1.1, 1.1, 3.003, 3.003, 7.007, 7.007},
		},
		{
			name: "random order",
			in:   []float32{10.1, math.MaxFloat32, 5.05, 100, 5.05, 42.5},
			want: []float32{5.05, 5.05, 10.1, 42.5, 100, math.MaxFloat32},
		},
		{
			name: "random order mixed",
			in:   []float32{10.1, -5.05, -20.1, 5.05, math.MaxFloat32, -5.05, 5.05, 2.02, 42.1, -math.MaxFloat32, 0},
			want: []float32{-math.MaxFloat32, -20.1, -5.05, -5.05, 0, 2.02, 5.05, 5.05, 10.1, 42.1, math.MaxFloat32},
		},
	}

	key := func(a float32) float32 {
		return a
	}

	sortFuncName := "Generic[float32](data, buf []float32)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]float32, len(tt.in))
			data := append([]float32{}, tt.in...)

			err := radixsort.Generic(data, buf, key)
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

func TestGenreicFloat64(t *testing.T) {
	tests := []struct {
		name string
		in   []float64
		want []float64
	}{
		{
			name: "empty slice",
			in:   []float64{},
			want: []float64{},
		},
		{
			name: "single element",
			in:   []float64{4.2},
			want: []float64{4.2},
		},
		{
			name: "single negative element",
			in:   []float64{-4.2},
			want: []float64{-4.2},
		},
		{
			name: "already sorted",
			in:   []float64{1.1, 2.02, 3.003, 4.4, 5.05},
			want: []float64{1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "already sorted negatives",
			in:   []float64{-5.05, -4.4, -3.003, -2.02, -1.1},
			want: []float64{-5.05, -4.4, -3.003, -2.02, -1.1},
		},
		{
			name: "already sorted mixed",
			in:   []float64{-5.05, -4.4, -3.003, -2.02, -1.1, 0, 1.1, 2.02, 3.003, 4.4, 5.05},
			want: []float64{-5.05, -4.4, -3.003, -2.02, -1.1, 0, 1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "reverse order",
			in:   []float64{5.05, 4.4, 3.003, 2.02, 1.1},
			want: []float64{1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "reverse order negatives",
			in:   []float64{-1.1, -2.02, -3.003, -4.4, -5.05},
			want: []float64{-5.05, -4.4, -3.003, -2.02, -1.1},
		},
		{
			name: "reverse order mixed",
			in:   []float64{-1.1, -2.02, -3.003, -4.4, -5.05, 5.05, 4.4, 3.003, 2.02, 1.1},
			want: []float64{-5.05, -4.4, -3.003, -2.02, -1.1, 1.1, 2.02, 3.003, 4.4, 5.05},
		},
		{
			name: "with duplicates",
			in:   []float64{7.007, 3.003, 7.007, 1.1, 3.003, 1.1},
			want: []float64{1.1, 1.1, 3.003, 3.003, 7.007, 7.007},
		},
		{
			name: "random order",
			in:   []float64{10.1, math.MaxFloat32, 5.05, 100, 5.05, 42.5},
			want: []float64{5.05, 5.05, 10.1, 42.5, 100, math.MaxFloat32},
		},
		{
			name: "random order mixed",
			in:   []float64{10.1, -5.05, -20.1, 5.05, math.MaxFloat32, -5.05, 5.05, 2.02, 42.1, -math.MaxFloat32, 0},
			want: []float64{-math.MaxFloat32, -20.1, -5.05, -5.05, 0, 2.02, 5.05, 5.05, 10.1, 42.1, math.MaxFloat32},
		},
	}

	key := func(a float64) float64 {
		return a
	}

	sortFuncName := "Generic[float64](data, buf []float64)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]float64, len(tt.in))
			data := append([]float64{}, tt.in...)

			err := radixsort.Generic(data, buf, key)
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

func TestGenericFloat64Random(t *testing.T) {
	size := 1_000_000
	input := make([]float64, size)
	for i := range input {
		input[i] = rand.Float64()
	}

	if slices.IsSorted(input) {
		t.Fatalf("terrible rand.rand")
	}

	data := append([]float64(nil), input...)
	buf := make([]float64, len(data))
	key := func(a float64) float64 {
		return a
	}

	err := radixsort.Generic(data, buf, key)
	if err != nil {
		t.Fatalf("%s failed: %v", "Generic[float64](randomData, buf, key)", err)
	}

	if !slices.IsSorted(data) {
		t.Errorf("%s failed to sort data correctly", "Generic[float64](randomData, buf, key)")
	}
}

func TestGenericStructsOfUint64(t *testing.T) {
	type person struct {
		Name string
		Age  uint64
	}

	tests := []struct {
		name string
		in   []person
		want []person
	}{
		{
			name: "empty slice",
			in:   []person{},
			want: []person{},
		},
		{
			name: "single element",
			in:   []person{{Name: "Zaphod", Age: 42}},
			want: []person{{Name: "Zaphod", Age: 42}},
		},
		{
			name: "already sorted",
			in: []person{
				{Name: "Baby Yoda", Age: 1},
				{Name: "Elon", Age: 2},
				{Name: "Gandalf", Age: 3},
				{Name: "Methuselah", Age: 4},
				{Name: "Dumbledore", Age: 5},
			},
			want: []person{
				{Name: "Baby Yoda", Age: 1},
				{Name: "Elon", Age: 2},
				{Name: "Gandalf", Age: 3},
				{Name: "Methuselah", Age: 4},
				{Name: "Dumbledore", Age: 5},
			},
		},
		{
			name: "reverse order",
			in: []person{
				{Name: "Dumbledore", Age: 5},
				{Name: "Methuselah", Age: 4},
				{Name: "Gandalf", Age: 3},
				{Name: "Elon", Age: 2},
				{Name: "Baby Yoda", Age: 1},
			},
			want: []person{
				{Name: "Baby Yoda", Age: 1},
				{Name: "Elon", Age: 2},
				{Name: "Gandalf", Age: 3},
				{Name: "Methuselah", Age: 4},
				{Name: "Dumbledore", Age: 5},
			},
		},
		{
			name: "with duplicates",
			in: []person{
				{Name: "TwinA", Age: 7},
				{Name: "CloneX", Age: 3},
				{Name: "TwinB", Age: 7},
				{Name: "Baby1", Age: 1},
				{Name: "CloneY", Age: 3},
				{Name: "Baby2", Age: 1},
			},
			want: []person{
				{Name: "Baby1", Age: 1},
				{Name: "Baby2", Age: 1},
				{Name: "CloneX", Age: 3},
				{Name: "CloneY", Age: 3},
				{Name: "TwinA", Age: 7},
				{Name: "TwinB", Age: 7},
			},
		},
		{
			name: "random order",
			in: []person{
				{Name: "TeenGoth", Age: 10},
				{Name: "TheAncientOne", Age: math.MaxUint64},
				{Name: "Kindergartener", Age: 5},
				{Name: "Centenarian", Age: 100},
				{Name: "AnotherKid", Age: 5},
				{Name: "Hitchhiker", Age: 42},
			},
			want: []person{
				{Name: "Kindergartener", Age: 5},
				{Name: "AnotherKid", Age: 5},
				{Name: "TeenGoth", Age: 10},
				{Name: "Hitchhiker", Age: 42},
				{Name: "Centenarian", Age: 100},
				{Name: "TheAncientOne", Age: math.MaxUint64},
			},
		},
	}

	sortByAge := func(a person) uint64 {
		return a.Age
	}

	sortFuncName := "Generic[person](data, buf []person)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]person, len(tt.in))
			data := append([]person{}, tt.in...)

			err := radixsort.Generic(data, buf, sortByAge)
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

func TestSortPointsByDistance(t *testing.T) {
	type point3D struct {
		X, Y, Z float64
	}

	tests := []struct {
		name string
		in   []point3D
		want []point3D
	}{
		{
			name: "empty slice",
			in:   []point3D{},
			want: []point3D{},
		},
		{
			name: "single point",
			in:   []point3D{{3, 4, 0}},
			want: []point3D{{3, 4, 0}},
		},
		{
			name: "already sorted by distance",
			in: []point3D{
				{0, 0, 0},
				{1, 0, 0},
				{1, 1, 0},
				{0, 0, 3},
			},
			want: []point3D{
				{0, 0, 0},
				{1, 0, 0},
				{1, 1, 0},
				{0, 0, 3},
			},
		},
		{
			name: "reverse order",
			in: []point3D{
				{0, 0, 3},
				{1, 1, 0},
				{1, 0, 0},
				{0, 0, 0},
			},
			want: []point3D{
				{0, 0, 0},
				{1, 0, 0},
				{1, 1, 0},
				{0, 0, 3},
			},
		},
		{
			name: "points with same distance (on sphere)",
			in: []point3D{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
				{-1, 0, 0},
			},
			want: []point3D{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
				{-1, 0, 0},
			},
		},
		{
			name: "mixed signs, same distances",
			in: []point3D{
				{-2, 0, 0},
				{0, -2, 0},
				{1, 1, 1},
				{0, 0, 1},
			},
			want: []point3D{
				{0, 0, 1},
				{1, 1, 1},
				{-2, 0, 0},
				{0, -2, 0},
			},
		},
		{
			name: "large distance",
			in: []point3D{
				{0, 0, 1},
				{math.MaxFloat64, 0, 0},
				{1, 0, 0},
			},
			want: []point3D{
				{0, 0, 1},
				{1, 0, 0},
				{math.MaxFloat64, 0, 0},
			},
		},
		{
			name: "random order",
			in: []point3D{
				{3, 0, 0},
				{0, 0, 0},
				{1, 1, 1},
				{0, -2, 0},
				{1, 2, 2},
			},
			want: []point3D{
				{0, 0, 0},
				{1, 1, 1},
				{0, -2, 0},
				{3, 0, 0},
				{1, 2, 2},
			},
		},
	}

	sortByDistanceFromCenter := func(p point3D) float64 {
		return p.X*p.X + p.Y*p.Y + p.Z*p.Z
	}

	sortFuncName := "Generic[point3D](data, buf []point3D)"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]point3D, len(tt.in))
			data := append([]point3D{}, tt.in...)

			err := radixsort.Generic(data, buf, sortByDistanceFromCenter)
			if err != nil {
				t.Fatalf("%s failed: %v", sortFuncName, err)
			}

			got := data

			if !cmp.Equal(tt.want, got) {
				t.Errorf("case: %s; SortPointsByDistance(%v) = %v, want %v", tt.name, tt.in, got, tt.want)
			}
		})
	}
}
