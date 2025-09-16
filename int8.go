package radixsort

import (
	"slices"
	"unsafe"
)

// Int8 sorts the given slice of int8 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than the data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Int8(data []int8, buf []uint8) {
	int8ver1call(data, buf)
}

func int8ver1call(data []int8, buf []uint8) {
	unsignedData := *(*[]uint8)(unsafe.Pointer(&data))
	radix8(unsignedData, buf)

	firstNegative := slices.IndexFunc(data, func(e int8) bool { return e < 0 })
	if firstNegative <= 0 {
		return
	}

	// Rearrange so that negative values come before positives.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)
}
