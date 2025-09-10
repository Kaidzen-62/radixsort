package radixsort

// Uint32 sorts the given slice of uint32 values in ascending order using the radix sort algorithm.
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// It panics if buf is shorter than data slice.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Uint32(data, buf []uint32) {
	radix32b8(data, buf)
}

// radix32b8 performs the internal radix sort implementation using 8-bit buckets.
// The buffer length must be at least as large as data.
func radix32b8(data, buf []uint32) {
	if len(buf) < len(data) {
		panic("Radixsort: buffer length is less than data length")
	}

	// offsets[d][b] stores prefix sums (insertion offsets) for digit d and offset b.
	// First they are used as frequency counters, then converted into offsets.
	offsets := [4][256]uint{}
	for _, v := range data {
		offsets[0][uint8(v>>(0*8))]++
		offsets[1][uint8(v>>(1*8))]++
		offsets[2][uint8(v>>(2*8))]++
		offsets[3][uint8(v>>(3*8))]++
	}

	// Calculate offsets.
	acc0 := offsets[0][0]
	acc1 := offsets[1][0]
	acc2 := offsets[2][0]
	acc3 := offsets[3][0]
	offsets[0][0] = 0
	offsets[1][0] = 0
	offsets[2][0] = 0
	offsets[3][0] = 0
	for i := 1; i < 256; i++ {
		offsets[0][i], acc0 = acc0, acc0+offsets[0][i]
		offsets[1][i], acc1 = acc1, acc1+offsets[1][i]
		offsets[2][i], acc2 = acc2, acc2+offsets[2][i]
		offsets[3][i], acc3 = acc3, acc3+offsets[3][i]
	}

	// Optimization: skip sorting passes where all elements in the digit are identical.
	uniqueOffsets := [4]uint{}
	for i := range 4 {
		// If all offsets are the same, there is only one unique value.
		if offsets[i][1] == offsets[i][255] {
			uniqueOffsets[i] = 1
			continue
		}

		for j := 1; j < 256; j++ {
			if offsets[i][j-1] != offsets[i][j] {
				uniqueOffsets[i]++
			}
			if offsets[i][j] == offsets[i][255] {
				break
			}
		}
	}

	swaps := 0
	src, dst := data, buf[:len(data)]
	for i := range 4 {
		if uniqueOffsets[i] < 2 {
			continue
		}
		swaps++

		for _, v := range src {
			index := offsets[i][uint8(v>>(i*8))]
			dst[index] = v
			offsets[i][uint8(v>>(i*8))]++
		}
		src, dst = dst, src
	}

	if swaps&1 == 1 {
		copy(data, src)
	}
}
