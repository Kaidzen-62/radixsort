package radixsort

import (
	"slices"
	"unsafe"
)

// Int64 sorts a slice of int64 values in ascending order.
//
// The data slice is sorted in place. The buf slice is used for temporary
// storage during sorting and must have len(buf) >= len(data).
//
// The buffer can be reused across multiple sort operations without clearing.
//
// Returns ErrInvalidBufferSize if len(buf) < len(data).
//
// Example:
//
//	data := []int64{-5, 2, -9, 1, 0}
//	buf := make([]uint64, len(data))
//	err := Int64(data, buf)
//	// data is now sorted: [-9, -5, 0, 1, 2]
func Int64(data []int64, buf []uint64) error {
	return int64ver1call(data, buf)
}

func int64ver1call(data []int64, buf []uint64) error {
	unsignedData := *(*[]uint64)(unsafe.Pointer(&data))
	err := radix64b8(unsignedData, buf)
	if err != nil {
		return err
	}

	firstNegative := slices.IndexFunc(data, func(e int64) bool { return e < 0 })
	if firstNegative <= 0 {
		return nil
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

	return nil
}
