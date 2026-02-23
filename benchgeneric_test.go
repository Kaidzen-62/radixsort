package radixsort_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/Kaidzen-62/radixsort"
)

type genericKeyFunc[E any, N radixsort.ConstraintNumbers] func(a E) N

type genericSortFunc[E any, N radixsort.ConstraintNumbers] func(data, buf []E, key func(a E) N) error

func BenchmarkGenericInt8(b *testing.B) {
	key := func(a int8) int8 {
		return a
	}

	benchmarkGenericSimple(b, radixsort.Generic, key, "GenericInt8")
}

func BenchmarkGenericInt64(b *testing.B) {
	key := func(a int64) int64 {
		return a
	}

	benchmarkGenericSimple(b, radixsort.Generic, key, "GenericInt64")
}

func BenchmarkGenericFloat64(b *testing.B) {
	key := func(a float64) float64 {
		return a
	}

	benchmarkGenericSimple(b, radixsort.Generic, key, "GenericFloat64")
}

type Point3 struct {
	X, Y, Z int64
}

func BenchmarkGenericStruct(b *testing.B) {
	key := func(p Point3) int64 {
		return p.Z
	}

	benchmarkGenericStruct(b, radixsort.Generic, key, "GenericPoint3")
}

func benchmarkGenericSimple[T radixsort.ConstraintNumbers](b *testing.B, sortFunc genericSortFunc[T, T], keyFunc genericKeyFunc[T, T], sortFuncName string) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("Radixsort%s_%d_%s", sortFuncName, size, mode)
			}(), func(b *testing.B) {
				data := generateData[T](size, mode)
				buf := make([]T, len(data))
				runtime.GC()

				b.ResetTimer()
				for b.Loop() {
					tmp := append([]T{}, data...)
					err := sortFunc(tmp, buf, keyFunc)
					if err != nil {
						b.Fatalf("%s failed: %v", sortFuncName, err)
					}
				}
			})
		}
	}
}

func benchmarkGenericStruct(b *testing.B, sortFunc genericSortFunc[Point3, int64], keyFunc genericKeyFunc[Point3, int64], sortFuncName string) {
	for _, size := range sizes {
		for _, mode := range modes {
			b.Run(func() string {
				return fmt.Sprintf("Radixsort%s_%d_%s", sortFuncName, size, mode)
			}(), func(b *testing.B) {
				data := generatePoint3Data(size, mode)
				buf := make([]Point3, len(data))
				runtime.GC()

				b.ResetTimer()
				for b.Loop() {
					tmp := append([]Point3{}, data...)
					err := sortFunc(tmp, buf, keyFunc)
					if err != nil {
						b.Fatalf("%s failed: %v", sortFuncName, err)
					}
				}
			})
		}
	}
}

func generatePoint3Data(n int, mode string) []Point3 {
	data := make([]Point3, n)

	switch mode {
	case "sorted":
		for i := range data {
			data[i] = Point3{X: int64(i), Y: int64(i), Z: int64(i)}
		}
	case "reverse":
		for i := range data {
			data[i] = Point3{X: int64(n - i), Y: int64(n - i), Z: int64(n - i)}
		}
	case "duplicates":
		for i := range data {
			data[i] = Point3{X: int64(i % 66), Y: int64(i % 66), Z: int64(i % 66)}
		}
	default: // "random"
		mu.Lock()
		for i := range n {
			rf := r.Float64()
			p := r.Intn(20)
			data[i] = Point3{
				X: int64(rf * float64(p)),
				Y: int64(rf * float64(p)),
				Z: int64(rf * float64(p)),
			}
		}
		mu.Unlock()
	}

	return data
}
