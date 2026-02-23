# Radix Sort in Go

This repository provides an implementation of the **Radix Sort** algorithm in Go.  
The package is designed for efficient, stable sorting of fixed-width integer types, with planned support for generic and custom data types.

---

## Overview

The `radixsort` package implements Radix Sort, a **non-comparative sorting algorithm**.  
Instead of directly comparing elements, the algorithm processes input by grouping keys based on individual digits (or bytes), starting from the least significant digit (LSD) to the most significant digit (MSD).

This implementation is optimized for unsigned and signed integers. Each sorting pass uses **counting sort on digit buckets**, ensuring **stability** of the algorithm. Optimizations are included to skip unnecessary passes when digits are uniform across the dataset.

---

## What is Radix Sort?

Radix Sort is a stable, integer-based sorting algorithm with time complexity:

- **O(n · k)**, where *n* is the number of elements and *k* is the number of digits (bytes) required to represent the maximum key.  
- For fixed-width integers (e.g., `uint16`, `uint32`, `uint64`), *k* is constant, which results in effective **O(n)** performance.

### Characteristics

- **Stable**: preserves the relative order of equal elements.  
- **Non-comparative**: does not rely on `<` or `>` operators.  
- **Digit-wise processing**: sorts values by successive digit positions, using counting sort as an internal 'subroutine'.  

---

## When to Use Radix Sort

Radix Sort is well-suited for:

- Large collections of integers with fixed width (`uint16`, `uint32`, `uint64`, `int32`, etc.).  
- Scenarios requiring **stable sorting**, where preserving order of equal keys is important.  
- Datasets where comparison-based algorithms (QuickSort, MergeSort) show overhead due to `O(n log n)` complexity.

---

## Features

- Optimized Radix Sort for unsigned integers (`uint16`, `uint32`, `uint64`).  
- Stable implementation using digit-based counting sort.  
- Skips redundant passes when all digit values are identical.  
- Planned support for:
  - Floating-point numbers
  - Strings
  - Generics and user-defined types

---

## Installation

To add the package to your Go project:

```bash
go get -u github.com/Kaidzen-62/radixsort
```

## Usage Example

```go
package main

import (
    "fmt"

    "github.com/Kaidzen-62/radixsort"
)

func main() {
    data := []uint16{170, 45, 75, 90, 802, 24, 2, 66}
    buf := make([]uint16, len(data))

    if err := radixsort.Uint16(data, buf); err != nil {
        panic(err)
    }

    fmt.Println(data) // Output: [2 24 45 66 75 90 170 802]
}
```

## Documentation

Complete reference is available at:
https://pkg.go.dev/github.com/Kaidzen-62/radixsort

## License

This project is licensed under the BSD 3-Clause License.
See the LICENSE file for details.

## TODO
- [ ] add interface version
- [x] add generic version
- [x] add generic benchmarks
- [ ] benchmart 'generic' version and optimize it
- [ ] add string version
- [ ] add examples and update TODO (generics) for godoc
- [ ] add comments (generics) for godoc

