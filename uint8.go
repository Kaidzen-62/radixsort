package radixsort

// Uint8 sorts the given slice of uint8 values in ascending order using the radix sort algorithm.
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// If the buffer length is invalid, it returns ErrInvalidBufferSize.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Uint8(data, buf []uint8) error {
	return radix8(data, buf)
}

// radix8b8 performs the internal radix sort implementation using 8-bit buckets.
// The buffer length must be at least as large as data.
func radix8(data, buf []uint8) error {
	if len(buf) < len(data) {
		return ErrInvalidBufferSize
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

	return nil
}
