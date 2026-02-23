package radixsort

import (
	"slices"
	"unsafe"
)

// Int8 sorts a slice of int8 values in ascending order.
//
// The data slice is sorted in place. The buf slice is used for temporary
// storage during sorting and must have len(buf) >= len(data).
//
// The buffer can be reused across multiple sort operations without clearing.
//
// Returns ErrInvalidBufferSize if len(buf) < len(data).
//
// See [Int64] for the 64-bit version and for usage example.
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
