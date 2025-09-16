package radixsort

import (
	"slices"
	"unsafe"
)

// Int16 sorts the given slice of int16 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than the data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Int16(data []int16, buf []uint16) {
	int16ver1call(data, buf)
}

func int16ver1call(data []int16, buf []uint16) {
	unsignedData := *(*[]uint16)(unsafe.Pointer(&data))
	radix16b8(unsignedData, buf)

	firstNegative := slices.IndexFunc(data, func(e int16) bool { return e < 0 })
	if firstNegative <= 0 {
		return
	}

	// Rearrange so that negative values come before positives.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)
}
