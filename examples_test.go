package radixsort_test

import (
	"fmt"

	"github.com/Kaidzen-62/radixsort"
)

func ExampleGeneric_float64() {
	data := []float64{3.14, -2.5, 0.0, 1.5, -1.0}
	buf := make([]float64, len(data))

	key := func(a float64) float64 {
		return a
	}

	if err := radixsort.Generic(data, buf, key); err != nil {
		panic(err)
	}
	fmt.Println(data)
	// Output:
	// [-2.5 -1 0 1.5 3.14]
}

// ExampleGeneric_struct demonstrates sorting a custom struct by an integer key.
func ExampleGeneric_struct() {
	type Product struct {
		ID    int
		Price int
	}

	data := []Product{
		{ID: 101, Price: 500},
		{ID: 102, Price: 150},
		{ID: 103, Price: 300},
	}
	buf := make([]Product, len(data))

	if err := radixsort.Generic(data, buf, func(p Product) int {
		return p.Price
	}); err != nil {
		panic(err)
	}

	for _, p := range data {
		fmt.Printf("ID: %d, Price: %d\n", p.ID, p.Price)
	}
	// Output:
	// ID: 102, Price: 150
	// ID: 103, Price: 300
	// ID: 101, Price: 500
}

func ExampleInt64() {
	data := []int64{-5, 3, -10, 0, 2}
	buf := make([]uint64, len(data))

	if err := radixsort.Int64(data, buf); err != nil {
		panic(err)
	}
	fmt.Println(data)
	// Output:
	// [-10 -5 0 2 3]
}

func ExampleUint64() {
	data := []uint64{170, 45, 75, 90, 802, 24, 2, 66}
	buf := make([]uint64, len(data))

	if err := radixsort.Uint64(data, buf); err != nil {
		panic(err)
	}
	fmt.Println(data)
	// Output:
	// [2 24 45 66 75 90 170 802]
}

// ExampleUint64_bufferReuse demonstrates that buffers can be reused
// across multiple sort operations without clearing them.
func ExampleUint64_bufferReuse() {
	// First sort
	data1 := []uint64{5, 2, 9, 1}
	buf := make([]uint64, len(data1))

	if err := radixsort.Uint64(data1, buf); err != nil {
		panic(err)
	}
	fmt.Println(data1)

	// Second sort - reuse the same buffer without clearing
	// Buffer contents after first sort: [1, 2, 5, 9]
	data2 := []uint64{8, 3, 7, 4}
	if err := radixsort.Uint64(data2, buf); err != nil {
		panic(err)
	}
	fmt.Println(data2)
	// Output:
	// [1 2 5 9]
	// [3 4 7 8]
}

// ExampleUint64_error demonstrates the error returned when the buffer
// is smaller than the data slice.
func ExampleUint64_error() {
	data := []uint64{3, 1, 2}
	buf := make([]uint64, 1) // buffer is too small
	err := radixsort.Uint64(data, buf)
	fmt.Println(err)
	// Output:
	// buffer length is less than data length
}
