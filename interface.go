package radixsort

import "github.com/sagernet/sing/common/x/constraints"

type Sortable[E constraints.Unsigned] interface {
	Len() int
	Get(i int) E
}
