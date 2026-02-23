package radixsort

import (
	"slices"
	"unsafe"
)

// Int32 sorts a slice of int32 values in ascending order.
//
// The data slice is sorted in place. The buf slice is used for temporary
// storage during sorting and must have len(buf) >= len(data).
//
// The buffer can be reused across multiple sort operations without clearing.
//
// Returns ErrInvalidBufferSize if len(buf) < len(data).
//
// See [Int64] for the 64-bit version and for usage example.
func Int32(data []int32, buf []uint32) error {
	return int32ver1call(data, buf)
}

func int32ver1call(data []int32, buf []uint32) error {
	unsignedData := *(*[]uint32)(unsafe.Pointer(&data))
	err := radix32b8(unsignedData, buf)
	if err != nil {
		return err
	}

	firstNegative := slices.IndexFunc(data, func(e int32) bool { return e < 0 })
	if firstNegative <= 0 {
		return nil
	}

	// Rearrange so that negative values come before positives.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)

	return nil
}
