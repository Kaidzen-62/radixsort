package archived

import (
	"unsafe"

	"github.com/Kaidzen-62/radixsort"
)

// Int64 sorts the given slice of int64 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than the data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Int64ver2calls(data []int64, buf []uint64) {
	p := positives(data, buf)
	n := negatives(data, buf[len(p):])

	unsignedData := *(*[]uint64)(unsafe.Pointer(&data))

	radixsort.Uint64(p, unsignedData)
	radixsort.Uint64(n, unsignedData)

	copy(unsignedData, n)
	copy(unsignedData[len(n):], p)
}

func positives(src []int64, buf []uint64) []uint64 {
	j := 0
	for i := range src {
		if src[i] >= 0 {
			buf[j] = uint64(src[i])
			j++
		}
	}

	return buf[:j]
}

func negatives(src []int64, buf []uint64) []uint64 {
	j := 0
	for i := range src {
		if src[i] < 0 {
			buf[j] = uint64(src[i])
			j++
		}
	}

	return buf[:j]
}
