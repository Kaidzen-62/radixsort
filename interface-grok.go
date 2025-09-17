package radixsort

// RadixSortable defines an interface for collections that can be sorted using radix sort.
// It allows indexed access for reading and writing values.
type RadixSortable interface {
	Len() int
	Get(i int) uint64
	Set(i int, val uint64)
	// NOTE: ниже я предложил добавить функцию которая вернет ссылку на слайс
}

// NOTE: набросок Грока меня удивил. Всё же это лучшая модель с которой я пока работал.
// но его решение не идеально и будет создавать кучу оверхэдов. Нужно проанализировать
// и подумать как можно это использовать.
// В общем как черновой вариант - мне понравилось, есть теперь над чем подумать.

// Uint64 sorts the given RadixSortable data in ascending order using radix sort.
// buf must be another RadixSortable of at least the same length as data.
// It panics if buf.Len() < data.Len().
func Uint_Grok(data, buf RadixSortable) {
	if buf.Len() < data.Len() {
		panic("Radixsort: buffer length is less than data length")
	}
	radix64b8_Grok(data, buf)
}

// radix64b8 performs the internal radix sort using 8-bit buckets (radix=256).
// NOTE: мы хотим эффективно использовать разрядности, если тип данных элемента массива меньше 64 бит.
// - Тут два может быть решения. Или выбираем нужную версию по разрядности. Но это полная жопа ведь придятся сделать еще 4 версии для uint и еще 4 для signed
// - Или можно сделать скипы циклов. Т.е. тип данных остается максимально 64bit, но лишние разряды не будут обрабатываться.
//
// NOTE: А ХОТЯ! Можно скомбинировать оба варианта. Да у нас будут лишние размерности (например [8][256]{} для 8bit чисел, лол) НО будут пропускатся циклы
// дженерики помогут заранее вычислить количество необходимых циклов!
// ПО ИТОГУ: получится интерфейс с дженериком (generic interface),
// data будет Интерфейсного типа, а buf будт generic типа.
func radix64b8_Grok(data, buf RadixSortable) {
	n := data.Len()
	// offsets[d][b] stores prefix sums for digit d (0-7 for 64-bit) and bucket b (0-255).
	offsets := [8][256]uint{}
	for i := range n {
		// NOTE: чтение значения элемента по которому сортируем
		v := data.Get(i)
		// NOTE: тут следует забэнчить что лучше: циклы или развернутый вариант
		for d := range 8 {
			offsets[d][uint8(v>>(d*8))]++
		}
	}

	// Convert counts to prefix sums.
	acc := [8]uint{}
	// NOTE: Вот это точно дичь. Если мы решим использовать второе решение, то все равно лучше развернуть цикл
	for d := range 8 {
		acc[d] = offsets[d][0]
		offsets[d][0] = 0
	}
	for b := 1; b < 256; b++ {
		// NOTE: А вот тут появляется пространство для маневра и действительно можно пропустить лишние циклы для лишних рязрядов
		// ХОТЯ: лучше этот момент забэнчить (цикл vs развернутый вариант)
		for d := range 8 {
			offsets[d][b], acc[d] = acc[d], acc[d]+offsets[d][b]
		}
	}

	// Optimization: count unique offsets per digit to skip uniform digits.
	// NOTE: тут тоже можно пропустить лишние циклы в зависимости от разрядности элемента
	uniqueOffsets := [8]uint{}
	for d := range 8 {
		if offsets[d][255] == 0 || offsets[d][1] == acc[d] {
			uniqueOffsets[d] = 1
			continue
		}
		prev := offsets[d][0]
		for b := 1; b < 256; b++ {
			if offsets[d][b] != prev {
				uniqueOffsets[d]++
				prev = offsets[d][b]
			}
			if offsets[d][b] == acc[d] {
				break
			}
		}
		if offsets[d][255] != acc[d] {
			uniqueOffsets[d]++
		}
	}

	swaps := 0
	src, dst := data, buf
	// NOTE: тут только можно поменять кол-во цилов. Если будем использовать дженерики, заранее вычислим кол-во байт
	for d := range 8 {
		if uniqueOffsets[d] < 2 {
			continue // Skip if all digits are identical.
		}
		swaps++

		for i := range n {
			// чтение самого элемента
			v := src.Get(i)
			bucket := uint8(v >> (d * 8))
			index := offsets[d][bucket]
			// NOTE: если это структура, то мы должны сортировать не числа из какого то поля,
			// a саму запись
			dst.Set(int(index), v)
			offsets[d][bucket]++
		}
		src, dst = dst, src
	}

	if swaps&1 == 1 {
		// If odd number of swaps, copy back to original data.
		// NOTE: вот тут хз. Я пологаю что copy все еще будет работать, при условии что мы добавим в интерфейс функцию которая вернет ссылку на оригинальный слайс
		for i := range n {
			data.Set(i, src.Get(i))
		}
	}
}

// Пример реализации интерфейса для слайса uint64 (чтобы протестировать).
type Uint64Slice []uint64

func (s Uint64Slice) Len() int              { return len(s) }
func (s Uint64Slice) Get(i int) uint64      { return s[i] }
func (s Uint64Slice) Set(i int, val uint64) { s[i] = val }

// Аналогично можно сделать для uint16, uint32 и т.д., с приведением типов в Get/Set.
type Uint16Slice []uint16

func (s Uint16Slice) Len() int              { return len(s) }
func (s Uint16Slice) Get(i int) uint64      { return uint64(s[i]) }
func (s Uint16Slice) Set(i int, val uint64) { s[i] = uint16(val) }

// Использование:
// data := Uint16Slice(myUint16Slice)
// buf := Uint16Slice(make([]uint16, len(myUint16Slice)))
// radixsort.Uint64(data, buf)  // Сортирует как uint64, но с приведением.ackage radixsort
