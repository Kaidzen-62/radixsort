package archived

// radix16b16 performs radix sort on the data slice using the provided buffer.
// The buffer must be at least as long as the data slice.
// This function modifies both the data and buffer slices.
//
// Researches showed that 16bit bucket is too big for stack
// and triggers compiler to shift bucket into heap.
// Also benchmarks showed that this version is slower than the 8bit version.
func radix16b16(data, buf []uint16) {
	if len(buf) < len(data) {
		panic("Radixsort: radix16b16: len(buf) < len(data)")
	}

	if len(data) <= 1 {
		return
	}

	// The bucket array is first used to count element frequencies,
	// and then reused to store offsets (prefix sums).
	bucket := [65536]uint{}
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

// radix16b8Opt performs radix sort on the data slice using the provided buffer.
// The buffer must be at least as long as the data slice.
// This function modifies both the data and buffer slices.
//
// This version includes an optimization that skips sorting if the input is already sorted.
// However, benchmarks show the check adds overhead for small to medium slices (under ~20K elements)
// and provides no measurable benefit even for large slices (over 100K), due to the cost of scanning
// and branch misprediction.
//
// Furthermore, it was decided that sorting functions should focus solely on sorting:
// pre-checks for sorted input shift responsibility away from the caller and add complexity
// without sufficient practical gain. Thus, this variant is archived for reference and benchmarking.
func radix16b8Opt(data, buf []uint16) {
	if len(buf) < len(data) {
		panic("Radixsort: radix16b8: len(buf) < len(data)")
	}

	if len(data) <= 1 {
		return
	}

	// The buckets are first used to count element frequencies,
	// and then reused to store offsets (prefix sums).
	prev := data[0]
	sorted := true

	bucket := [2][256]uint{}
	for _, v := range data {
		bucket[0][uint8(v>>(0*8))]++
		bucket[1][uint8(v>>(1*8))]++

		sorted = sorted && (prev <= v)
		prev = v
	}
	if sorted {
		return
	}

	// Calculate offsets.
	offsets := [2][256]uint{}
	for i := 1; i < 256; i++ {
		offsets[0][i] = offsets[0][i-1] + bucket[0][i-1]
		offsets[1][i] = offsets[1][i-1] + bucket[1][i-1]
	}

	src, dst := data, buf[:len(data)]
	for i := range 2 {
		for _, v := range src {
			index := offsets[i][uint8(v>>(i*8))]
			dst[index] = v
			offsets[i][uint8(v>>(i*8))]++
		}
		src, dst = dst, src
	}

	copy(data, src)
}
