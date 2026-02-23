// Package radixsort provides high-performance radix sort implementations
// for fixed-width integer and floating-point types.
//
// # Overview
//
// Radix sort is a non-comparative sorting algorithm that processes elements
// by grouping them based on individual digits or bytes, from least significant
// to most significant. This implementation achieves O(n) time complexity for
// fixed-width types, making it faster than comparison-based algorithms for
// large datasets.
//
// # Features
//
//   - Stable sorting: preserves the relative order of equal elements
//   - In-place sorting with a temporary buffer
//   - Optimized for unsigned integers (uint8, uint16, uint32, uint64)
//   - Support for signed integers (int8, int16, int32, int64)
//   - Generic sorting for custom types with numeric keys
//   - Automatic skip of redundant sorting passes
//
// # Usage
//
// Each sorting function requires a temporary buffer of the same length as
// the input data. This design avoids allocations during sorting and allows
// buffer reuse across multiple sort operations.
//
// Basic example for uint64:
//
//	data := []uint64{5, 2, 9, 1, 5, 6}
//	buf := make([]uint64, len(data))
//	err := radixsort.Uint64(data, buf)
//	if err != nil {
//	    // handle error
//	}
//
// For signed integers, the algorithm handles the sign bit automatically:
//
//	data := []int64{-5, 2, -9, 1, 0}
//	buf := make([]uint64, len(data))
//	err := radixsort.Int64(data, buf)
//
// # Generic Sorting
//
// The [Generic] function allows sorting custom types by extracting a numeric key:
//
//	type Item struct {
//	    ID   int
//	    Score float64
//	}
//
//	items := []Item{{1, 95.5}, {2, 87.3}, {3, 92.1}}
//	buf := make([]Item, len(items))
//	err := radixsort.Generic(items, buf, func(i Item) float64 { return i.Score })
//
// # Performance Considerations
//
// Radix sort excels when:
//   - Sorting large collections of fixed-width numbers
//   - Stable ordering is required
//
// For small slices (< 1000 elements), standard library sort may be faster
// due to lower constant factors.
//
// # Error Handling
//
// All functions return [ErrInvalidBufferSize] if the buffer is smaller than
// the data slice. Ensure len(buf) >= len(data) before calling.
package radixsort
