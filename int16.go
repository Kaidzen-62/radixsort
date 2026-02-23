package radixsort

import (
	"slices"
	"unsafe"
)

// Int16 sorts a slice of int16 values in ascending order.
//
// The data slice is sorted in place. The buf slice is used for temporary
// storage during sorting and must have len(buf) >= len(data).
//
// The buffer can be reused across multiple sort operations without clearing.
//
// Returns ErrInvalidBufferSize if len(buf) < len(data).
//
// See [Int64] for the 64-bit version and for usage example.
func Int16(data []int16, buf []uint16) error {
	return int16ver1call(data, buf)
}

func int16ver1call(data []int16, buf []uint16) error {
	unsignedData := *(*[]uint16)(unsafe.Pointer(&data))
	err := radix16b8(unsignedData, buf)
	if err != nil {
		return err
	}

	firstNegative := slices.IndexFunc(data, func(e int16) bool { return e < 0 })
	if firstNegative <= 0 {
		return nil
	}

	// Rearrange so that negative values come before positives.
	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)

	return nil
}
