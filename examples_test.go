package radixsort_test

import (
	"fmt"

	"github.com/Kaidzen-62/radixsort"
)

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

func ExampleUint64_error() {
	data := []uint64{3, 1, 2}
	buf := make([]uint64, 1) // buffer is too small
	err := radixsort.Uint64(data, buf)
	fmt.Println(err)
	// Output:
	// buffer length is less than data length
}
