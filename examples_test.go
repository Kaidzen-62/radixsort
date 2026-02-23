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

func ExampleGeneric_struct() {
	type Person struct {
		Name string
		Age  uint64
	}

	data := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Diana", 25},
	}
	buf := make([]Person, len(data))

	key := func(p Person) uint64 {
		return p.Age
	}

	if err := radixsort.Generic(data, buf, key); err != nil {
		panic(err)
	}
	for _, p := range data {
		fmt.Printf("%s: %d\n", p.Name, p.Age)
	}
	// Output:
	// Bob: 25
	// Diana: 25
	// Alice: 30
	// Charlie: 35
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
