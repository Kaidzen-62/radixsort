package radixsort

import (
	"slices"
	"unsafe"
)

func Int64(data, buf []int64) {
}

func int64ver2calls(data []int64, buf []uint64) {
	p := positives(data, buf)
	n := negatives(data, buf[len(p):])

	unsignedData := *(*[]uint64)(unsafe.Pointer(&data))

	radix64b8(p, unsignedData)
	radix64b8(n, unsignedData)

	copy(unsignedData, n)
	copy(unsignedData[len(n):], p)
}

func int64ver1call(data []int64, buf []uint64) {
	unsignedData := *(*[]uint64)(unsafe.Pointer(&data))

	radix64b8(unsignedData, buf)

	firstNegative := slices.IndexFunc(data, func(e int64) bool { return e < 0 })
	if firstNegative <= 0 {
		return
	}

	copy(buf, unsignedData)
	copy(unsignedData, buf[firstNegative:])
	copy(unsignedData[len(data)-firstNegative:], buf)
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
