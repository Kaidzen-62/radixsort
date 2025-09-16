package radixsort

import (
	"slices"
	"unsafe"
)

// Int64 sorts the given slice of int64 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than the data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Int64(data []int64, buf []uint64) {
	int64ver1call(data, buf)
}

func int64ver1call(data []int64, buf []uint64) {
	unsignedData := *(*[]uint64)(unsafe.Pointer(&data))
	radix64b8(unsignedData, buf)

	firstNegative := slices.IndexFunc(data, func(e int64) bool { return e < 0 })
	if firstNegative <= 0 {
		return
	}

	// After sorting unsigned values:
	//   - Positive numbers are placed before negative numbers.
	// To restore correct signed order:
	//   - Copy the whole sorted array into the buffer.
	//   - Move the block of negative numbers to the front.
	//   - Append the block of positive numbers after them.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)
}
