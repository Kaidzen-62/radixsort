package radixsort

import (
	"slices"
	"unsafe"
)

// Int32 sorts the given slice of int32 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than the data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Int32(data []int32, buf []uint32) {
	int32ver1call(data, buf)
}

func int32ver1call(data []int32, buf []uint32) {
	unsignedData := *(*[]uint32)(unsafe.Pointer(&data))
	radix32b8(unsignedData, buf)

	firstNegative := slices.IndexFunc(data, func(e int32) bool { return e < 0 })
	if firstNegative <= 0 {
		return
	}

	// Rearrange so that negative values come before positives.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)
}
