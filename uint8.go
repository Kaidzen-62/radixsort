package radixsort

// Uint8 sorts the data slice of type uint8 using radix sort algorithm with a temporary buffer.
// It panics if buf is shorter than data slice.
func Uint8(data, buf []uint8) {
	radix8(data, buf)
}

// radix8 performs radix sort on the data slice using the provided buffer.
// The buffer must be at least as long as the data slice.
// This function modifies both the data and buffer slices.
func radix8(data, buf []uint8) {
	if len(buf) < len(data) {
		panic("Radixsort: radix8: len(buf) < len(data)")
	}

	if len(data) <= 1 {
		return
	}

	// The bucket array is first used to count element frequencies,
	// and then reused to store offsets (prefix sums).
	bucket := [256]uint{}
	for _, v := range data {
		bucket[v]++
	}

	// Calculate offsets.
	acc := bucket[0]
	bucket[0] = 0
	for i := 1; i < len(bucket); i++ {
		bucket[i], acc = acc, acc+bucket[i]
	}

	out := buf[:len(data)]
	for _, v := range data {
		index := bucket[v]
		out[index] = v
		bucket[v]++
	}

	copy(data, out)
}
