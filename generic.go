package radixsort

import (
	"fmt"
	"unsafe"

	"github.com/sagernet/sing/common/x/constraints"
)

type ConstraintNumbers interface {
	constraints.Integer | constraints.Float
}

// Generic sorts a slice of elements using radix sort with a custom key extractor.
//
// Generic allows sorting any slice type E by extracting a numeric key N
// (integer or float type) from each element.
//
// Parameters:
//   - data: the slice to sort (modified in place)
//   - buf: temporary buffer, must have len(buf) >= len(data)
//   - key: function that extracts a numeric sort key from each element
//
// The key function is called once per element per sorting pass. For best
// performance, keep the key extraction simple and fast.
//
// The buffer can be reused across multiple sort operations without clearing.
//
// Returns ErrInvalidBufferSize if len(buf) < len(data).
//
// Example with float64 keys:
//
//	type Item struct{ Score float64 }
//	items := []Item{{95.5}, {87.3}, {92.1}}
//	buf := make([]Item, len(items))
//	err := Generic(items, buf, func(i Item) float64 { return i.Score })
//
// Example with int keys:
//
//	type User struct{ ID int }
//	users := []User{{3}, {1}, {2}}
//	buf := make([]User, len(users))
//	err := Generic(users, buf, func(u User) int { return u.ID })
func Generic[E any, N ConstraintNumbers](data, buf []E, key func(a E) N) error {
	if len(data) < 2 {
		return nil
	}

	if len(buf) < len(data) {
		return fmt.Errorf("buffer length is less than data length")
	}

	var keyZeroValue N
	sizeofKey := unsafe.Sizeof(keyZeroValue)

	var signBit uint64 = 1 << (sizeofKey*8 - 1)

	// default for unsigned integers
	unsignedKey := func(a E) uint64 {
		v := key(a)
		return *(*uint64)(unsafe.Pointer(&v))
	}

	// For signed types:
	/*
		Array of: 2 1 0 -1 -2
			0x02 0x01 0x00 0xff 0xfe
		becomes:
			0x82 0x81 0x80 0x7f 0x7e
		and will sort as:
			0x7e 0x7f 0x80 0x81 0x82
	*/
	switch any(keyZeroValue).(type) {
	case int, int8, int16, int32, int64:
		unsignedKey = func(a E) uint64 {
			v := key(a)
			uv := *(*uint64)(unsafe.Pointer(&v))
			return uv ^ signBit
		}
	case float32, float64:
		unsignedKey = func(a E) uint64 {
			v := key(a)
			uv := *(*uint64)(unsafe.Pointer(&v))

			if v < 0 {
				return (^uv + 1) & (signBit - 1)
			}
			return uv ^ signBit
		}
	}

	// offsets[d][b] stores prefix sums (insertion offsets) for digit d and offsets b.
	// First they are used as frequency counters, then converted into offsets.
	offsets := [8][256]uint{}
	for _, e := range data {
		// NOTE: тут следует забэнчить что лучше: циклы или развернутый вариант
		for d := range sizeofKey {
			b := byte(unsignedKey(e) >> (d * 8))
			offsets[d][b]++
		}
	}

	// Convert counts into prefix sums (offsets).
	// NOTE: стоит забенчить что лучше: развернутный или свернутый цикл
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

	// for b := 1; b < 256; b++ {
	//   for d := range sizeofKey {
	//     NOTE: А вот тут появляется пространство для маневра и действительно можно пропустить лишние циклы для лишних рязрядов
	//     ХОТЯ: лучше этот момент забэнчить (цикл vs развернутый вариант)
	//     offsets[d][b], acc[d] = acc[d], acc[d]+offsets[d][b]
	//   }
	// }

	// Optimization: skip sorting passes where all elements in the digit are identical.
	uniqueOffsets := [8]uint{}
	for i := range sizeofKey {
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
	for d := range sizeofKey {
		if uniqueOffsets[d] < 2 {
			continue
		}
		swaps++

		for _, e := range src {
			b := byte(unsignedKey(e) >> (d * 8))
			index := offsets[d][b]
			dst[index] = e
			offsets[d][b]++
		}
		src, dst = dst, src
	}

	if swaps&1 == 1 {
		copy(data, src)
	}

	return nil
}
