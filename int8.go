package radixsort

import (
	"slices"
	"unsafe"
)

// Int8 sorts the given slice of int8 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// If the buffer length is invalid, it returns ErrInvalidBufferSize.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Int8(data []int8, buf []uint8) error {
	return int8ver1call(data, buf)
}

func int8ver1call(data []int8, buf []uint8) error {
	unsignedData := *(*[]uint8)(unsafe.Pointer(&data))
	err := radix8(unsignedData, buf)
	if err != nil {
		return err
	}

	firstNegative := slices.IndexFunc(data, func(e int8) bool { return e < 0 })
	if firstNegative <= 0 {
		return nil
	}

	// Rearrange so that negative values come before positives.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)

	return nil
}
