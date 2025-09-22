package radixsort

// Uint64 sorts the given slice of uint64 values in ascending order using the radix sort algorithm.
//
// A temporary buffer (buf) is required, and its length must be at least as large as data.
// If the buffer length is invalid, it returns ErrInvalidBufferSize.
//
// Both data and buf will be modified during sorting.
// The algorithm is stable and runs in O(n) time complexity.
func Uint64(data, buf []uint64) error {
	return radix64b8(data, buf)
}

// radix64b8 performs the internal radix sort implementation using 8-bit buckets.
// The buffer length must be at least as large as data.
func radix64b8(data, buf []uint64) error {
	if len(buf) < len(data) {
		return ErrInvalidBufferSize
	}

	// offsets[d][b] stores prefix sums (insertion offsets) for digit d and offsets b.
	// First they are used as frequency counters, then converted into offsets.
	offsets := [8][256]uint{}
	for _, v := range data {
		offsets[0][uint8(v>>(0*8))]++
		offsets[1][uint8(v>>(1*8))]++
		offsets[2][uint8(v>>(2*8))]++
		offsets[3][uint8(v>>(3*8))]++
		offsets[4][uint8(v>>(4*8))]++
		offsets[5][uint8(v>>(5*8))]++
		offsets[6][uint8(v>>(6*8))]++
		offsets[7][uint8(v>>(7*8))]++
	}

	// Convert counts into prefix sums (offsets).
	acc := [8]uint{
		offsets[0][0],
		offsets[1][0],
		offsets[2][0],
		offsets[3][0],
		offsets[4][0],
		offsets[5][0],
		offsets[6][0],
		offsets[7][0],
	}
	offsets[0][0] = 0
	offsets[1][0] = 0
	offsets[2][0] = 0
	offsets[3][0] = 0
	offsets[4][0] = 0
	offsets[5][0] = 0
	offsets[6][0] = 0
	offsets[7][0] = 0
	for i := 1; i < 256; i++ {
		offsets[0][i], acc[0] = acc[0], acc[0]+offsets[0][i]
		offsets[1][i], acc[1] = acc[1], acc[1]+offsets[1][i]
		offsets[2][i], acc[2] = acc[2], acc[2]+offsets[2][i]
		offsets[3][i], acc[3] = acc[3], acc[3]+offsets[3][i]
		offsets[4][i], acc[4] = acc[4], acc[4]+offsets[4][i]
		offsets[5][i], acc[5] = acc[5], acc[5]+offsets[5][i]
		offsets[6][i], acc[6] = acc[6], acc[6]+offsets[6][i]
		offsets[7][i], acc[7] = acc[7], acc[7]+offsets[7][i]
	}

	// Optimization: skip sorting passes where all elements in the digit are identical.
	//
	// Normally this is done by checking how many non-zero counters exist.
	// If only one counter is non-zero, then all values in this digit are the same
	// and sorting by this byte can be skipped.
	//
	// Here we use offsets instead of counters. For example:
	// If offsets[i][1] == offsets[i][255], it means all elements for digit i are 0,
	// and this sorting pass can be skipped entirely.
	//
	// Additionally, we count how many times the offset changes (uniqueOffsets).
	// This prepares the ground for a future optimization:
	// if there are exactly 2 distinct values in a digit (e.g. only 2 and 3 appear),
	// we might be able to sort it more efficiently.
	//
	// In practice, the speedup is noticeable only on arrays of size < 100_000 elements;
	// on larger arrays, the benefit becomes insignificant :-(
	uniqueOffsets := [8]uint{}
	for i := range 8 {
		if offsets[i][255] == 0 || offsets[i][1] == acc[i] {
			uniqueOffsets[i] = 1
			continue
		}

		for j := 1; j < 256; j++ {
			if offsets[i][j] != offsets[i][j-1] {
				uniqueOffsets[i]++
			}

			if offsets[i][j] == acc[i] {
				break
			}
		}

		if offsets[i][255] != acc[i] {
			uniqueOffsets[i]++
		}
	}

	swaps := 0
	src, dst := data, buf[:len(data)]
	for i := range 8 {
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

	return nil
}
