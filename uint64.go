package radixsort

// Uint64 sorts the data slice of type uint64 using radix sort algorithm with a temporary buffer.
// It panics if buf is shorter than data slice.
func Uint64(data, buf []uint64) {
	radix64b8(data, buf)
}

// radix64b8 performs radix sort on the data slice using the provided buffer.
// The buffer must be at least as long as the data slice.
// This function modifies both the data and buffer slices.
func radix64b8(data, buf []uint64) {
	if len(buf) < len(data) {
		panic("Radixsort: radix32b8: len(buf) < len(data)")
	}

	// The buckets are first used to count element frequencies,
	// and then reused to store offsets (prefix sums).
	bucket := [8][256]uint{}
	for _, v := range data {
		bucket[0][uint8(v>>(0*8))]++
		bucket[1][uint8(v>>(1*8))]++
		bucket[2][uint8(v>>(2*8))]++
		bucket[3][uint8(v>>(3*8))]++
		bucket[4][uint8(v>>(4*8))]++
		bucket[5][uint8(v>>(5*8))]++
		bucket[6][uint8(v>>(6*8))]++
		bucket[7][uint8(v>>(7*8))]++
	}

	// Calculate offsets.
	acc0 := bucket[0][0]
	acc1 := bucket[1][0]
	acc2 := bucket[2][0]
	acc3 := bucket[3][0]
	acc4 := bucket[4][0]
	acc5 := bucket[5][0]
	acc6 := bucket[6][0]
	acc7 := bucket[7][0]
	bucket[0][0] = 0
	bucket[1][0] = 0
	bucket[2][0] = 0
	bucket[3][0] = 0
	bucket[4][0] = 0
	bucket[5][0] = 0
	bucket[6][0] = 0
	bucket[7][0] = 0
	for i := 1; i < 256; i++ {
		bucket[0][i], acc0 = acc0, acc0+bucket[0][i]
		bucket[1][i], acc1 = acc1, acc1+bucket[1][i]
		bucket[2][i], acc2 = acc2, acc2+bucket[2][i]
		bucket[3][i], acc3 = acc3, acc3+bucket[3][i]
		bucket[4][i], acc4 = acc4, acc4+bucket[4][i]
		bucket[5][i], acc5 = acc5, acc5+bucket[5][i]
		bucket[6][i], acc6 = acc6, acc6+bucket[6][i]
		bucket[7][i], acc7 = acc7, acc7+bucket[7][i]
	}

	src, dst := data, buf[:len(data)]
	for i := range 8 {
		for _, v := range src {
			index := bucket[i][uint8(v>>(i*8))]
			dst[index] = v
			bucket[i][uint8(v>>(i*8))]++
		}
		src, dst = dst, src
	}

	copy(data, src)
}
