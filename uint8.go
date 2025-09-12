package radixsort

// Uint8 sorts the given slice of uint8 values in ascending order using the radix sort algorithm.
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Uint8(data, buf []uint8) {
	radix8(data, buf)
}

// radix8b8 performs the internal radix sort implementation using 8-bit buckets.
// The buffer length must be at least as large as data.
func radix8(data, buf []uint8) {
	if len(buf) < len(data) {
		panic("Radixsort: buffer length is less than data length")
	}

	// offsets[d][b] stores prefix sums (insertion offsets) for digit d and offsets b.
	// First they are used as frequency counters, then converted into offsets.
	offsets := [256]uint{}
	for _, v := range data {
		offsets[v]++
	}

	// Calculate offsets.
	acc := offsets[0]
	offsets[0] = 0
	for i := 1; i < len(offsets); i++ {
		offsets[i], acc = acc, acc+offsets[i]
	}

	dst := buf[:len(data)]
	for _, v := range data {
		index := offsets[v]
		dst[index] = v
		offsets[v]++
	}

	copy(data, dst)
}
