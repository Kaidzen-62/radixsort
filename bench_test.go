package radixsort_test

import (
	"math/rand"
	"reflect"
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

var sizes = []int{100, 1000, 10_000, 100_000, 1_000_000}
var modes = []string{"random", "sorted", "reverse", "duplicates"}

// Key for cache map
type dataKey struct {
	n    int
	mode string
	typ  string // for distinguishing types, e.g. "int", "int64", etc.
}

// Global cache and mutex
var (
	cache = make(map[dataKey][]any)
	mu    sync.RWMutex
	r     = rand.New(rand.NewSource(time.Now().UnixNano())) // local random generator
)

func generateData[T constraints.Integer](n int, mode string) []T {
	key := dataKey{n: n, mode: mode, typ: typeof[T]()}

	mu.RLock()
	if cached, found := cache[key]; found {
		mu.RUnlock()
		// Convert back to the required type
		res := make([]T, len(cached))
		for i, v := range cached {
			res[i] = v.(T)
		}
		return res
	}
	mu.RUnlock()

	// Generate data
	data := make([]T, n)
	switch mode {
	case "sorted":
		for i := range data {
			data[i] = T(i)
		}
	case "reverse":
		for i := range data {
			data[i] = T(n - i)
		}
	case "duplicates":
		for i := range data {
			data[i] = T(i % 66)
		}
	default: // "random"
		mu.Lock()
		for i := range data {
			data[i] = T(r.Intn(n * 10))
		}
		mu.Unlock()
	}

	// Save to cache
	mu.Lock()
	cache[key] = make([]any, len(data))
	for i, v := range data {
		cache[key][i] = v
	}
	mu.Unlock()

	return data
}

// typeof — helper function to get type name
func typeof[T any]() string {
	var zero T
	return reflect.TypeOf(zero).String()
}
