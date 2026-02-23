package radixsort

import "errors"

// ErrInvalidBufferSize is returned when the provided buffer slice is smaller
// than the data slice to be sorted.
//
// All sorting functions in this package require a temporary buffer for
// intermediate storage during the radix sort process. The buffer must have
// a length greater than or equal to the data slice:
//
//	len(buf) >= len(data)
//
// This error indicates that the buffer is too small and the sorting
// operation cannot proceed safely.
//
// Example:
//
//	data := []uint64{5, 2, 9}
//	buf := make([]uint64, 2) // too small!
//	err := radixsort.Uint64(data, buf)
//	// err == radixsort.ErrInvalidBufferSize
var ErrInvalidBufferSize = errors.New("buffer length is less than data length")
