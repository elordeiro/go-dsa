# Iters Package

> `Go Docs`  
> [![Go Reference](https://pkg.go.dev/badge/github.com/elordeiro/go/iters.svg)](https://pkg.go.dev/github.com/elordeiro/go/iters)

## Overview

The `iters` package provides a collection of iterators designed to enhance the traversal and transformation of data structures in Go. Leveraging the new iterator support introduced in Go 1.23, this package offers a range of utilities to simplify working with collections.

## Features

-   **Lazy Evaluation:** Efficient processing of elements with lazy evaluation. Process elements only when needed.
-   **Chainable Operations:** Support for chaining multiple iterator operations for complex data transformations.
-   **Range-Based Iteration:** Simplified range-based loops for common data structures.

## Installation

To install the `iters` package, run:

```bash
go get github.com/elordeiro/go/iters
```

Import the package in your Go code:

```go
import "github.com/elordeiro/go/iters"
```

or

```go
import it "github.com/elordeiro/go/iters"
```

## Usage

-   Most of the functions in this package have a `func()` and `func2()` version.
-   The `func()` version will accept any iterable that implements the `Iterable` interface.
    -   The `Iterable` interface requires the implementation of the `Len()`, `Values()`, and `All()` methods.
    -   This package implements the `Iterable` interface for `Seq` and `ItSlice`.
-   The `func2()` version will accept any iterable that implements the `Iterable2` interface.
    -   The `Iterable2` interface requires the implementation of the `Len()`, `Keys()`, `Values()`, and `All()` methods.
    -   This package implements the `Iterable2` interface for `Seq2` and `ItMap`.
-   For iterators that need to return more than two values, the `Seq2` type is used with a `Pair` struct.
-   If an iterator is not yet created, you can pass in a slice or a map to `Iterable()` or `Iterable2()` respectively to create one.

## Iterator Methods:

-   **Len()**: Len returns the length of the iterator.
-   **Keys()**: Keys returns a `Seq` iterator from the keys of a `Seq2` or `ItMap`.
-   **Values()**: Values returns a `Seq` iterator from the values of a `Seq`, `Seq2`, `ItSlice`, or `ItMap`.
-   **All()**:
    -   From a `Seq` or `ItSlice`, All returns an index-value `Seq2` iterator (`i` starts at `0`).
    -   From a `Seq2` or `ItMap`, All returns a key-value `Seq2` iterator.
-   **Sink()**: Sink converts a `Seq` to `iter.Seq` to be used with the standard library.
-   **String()**: String returns a string representation of the iterator.

For converting standard library iterators to be used within this package use:

-   **Surf()**: Surf function is used to convert an `iter.Seq` to `Seq` to be used within this package.
-   **Surf2()**: Surf2 function is used to convert an `iter.Seq2` to `Seq2` to be used within this package.

## Iterator Functions:

### The following functions only have a `func()` version.

-   **Range()**: Range returns an iterator over a range of values. It can take either one, two or three arguments.
    -   1 arg: value is the end value and the start value is `0`.
    -   2 args: values are the start and end values.
    -   3 args: values are the start, end, and step values.
    -   Range automatically determines the direction of the range, and panics if it detects an infinite loop.
-   **Count()**: Count returns an iterator that counts up from a start value.
-   **CountDown()**: CountDown returns an iterator that counts down from a start value.

### The following functions have a `func2()` version.

> **_Itertools_**:

-   **Enumerate()**: Enumerate returns a `Seq2` iterator with an index attached to each value starting from a given index.
    -   For a `Seq` or `ItSlice`, it returns an index-value iterator.
    -   For a `Seq2` or `ItMap`, it returns an index-pair iterator.
-   **Zip()**: Zip returns a `Seq2` iterator with elements from two iterators.
    -   For 2 `Seq`s or `ItSlice`s, it returns a value-value iterator.
    -   For 2 `Seq2`s or `ItMap`s, it returns a pair-pair iterator.
-   **ForEach()**: ForEach consumes the entire iterator and calls a function with each element.
-   **Repeat()**: Repeat returns an iterator that repeats a value indefinitely.
-   **Cycle()**: Cycles returns an iterator that repeats the elements of the given iterator indefinitely.
-   **Chain()**: Chain returns an iterator that chains multiple iterators together.
-   **Backwards()**: Backwards returns an iterator that iterates over the elements in reverse order.
-   **Take()**: Take returns the first n elements of the iterator.
-   **Drop()**: Drop returns an iterator that skips the first n elements.
-   **TakeBetween()**: TakeBetween returns an iterator that takes elements between two indices (inclusive).
-   **Rotate()**: Rotate returns an iterator that rotates the elements of the iterator, shifting them to the left or right.

> **_Functional_**:

-   **Filter()**: Filter returns an iterator whose elements satisfy the predicate.
-   **Map()**: Map returns an iterator that applies a function to each element.
-   **Reduce()**: Reduce returns a single value by applying a function to each element.
-   **TakeWhile()**: TakeWhile returns an iterator that takes elements from the iterator while the predicate is true.
-   **DropWhile()**: DropWhile returns an iterator that drops elements from the iterator while the predicate is true.
-   **With()**: With returns an iterator that calls a function with each element.
-   **OnEmpty()**: Else returns an iterator that calls a function if the iterator is exhausted.

### Convenience Functions:

> **_Seq2_** functions:

-   **Split()**: Splits takes an `Iterable[Pair[K, V]]` and returns a `Seq2[K, V]`.
-   **SwapKV()**: Swaps the key and value of a `Seq2` iterator.

> **_Seq_** functions:

-   **Sum()**: Sum returns the sum of the elements in the iterator.
-   **Product()**: Product returns the product of the elements in the iterator.
-   **Min()**: Min returns the minimum element in the iterator.
-   **Max()**: Max returns the maximum element in the iterator.
-   **All()**: All returns true if all elements in the iterator satisfy the predicate.
-   **Any()**: Any returns true if any element in the iterator satisfies the predicate.
-   **None()**: None returns true if no element in the iterator satisfies the predicate.

## Examples

### Basic iterators

Here's are some examples of how to use the `iters` package:

```go
package main

import (
    "fmt"
    it "github.com/elordeiro/go/iters"
)

func main() {
    for v := range it.Range(1, 6) {
        fmt.Print(v, " ")
    }
    fmt.Println()
    // Output: 1 2 3 4 5

    slice := it.Iterable([]bool{true, false, true, false, true})
    fmt.Println(it.Any(slice.Values()))
    // Output: true

    result := ""
    i := 31
    Else := func() { result = fmt.Sprint(i, " is prime") }
    for range it.OrElse(it.Range(2, i), Else) {
        if i%2 == 0 {
            result = fmt.Sprint(i, " is not prime")
            break
        }
    }
    fmt.Println(result)
    // Output: 31 is prime
}
```

### Chaining operations

You can chain multiple operations using the iterators:

```go
package main

import (
    "fmt"
    it "github.com/elordeiro/go/iters"
)

func main() {
    square := func(x int) int { return x * x }
    for v := range it.Map(it.Range(6), square) {
        fmt.Print(v, " ")
    }
    fmt.Println()
    // Output: 1 4 9 16 25

    sum := func(a, b int) int { return a + b }
    result := it.Reduce(it.Range(1, 6), sum)
    fmt.Println(result)
    // Output: 15

    for v := range it.Take(9, it.Cycle(it.Range(0, 3))) {
        fmt.Print(v, " ")
    }
    fmt.Println()
    // Output: 0 1 2 0 1 2 0 1 2

    for v := range it.Chain(it.Range(2), it.Range(2, 5), it.Range(5, 10, 2)) {
        fmt.Print(v, " ")
    }
    fmt.Println()
    // Output: 0 1 2 3 4 5 7 9
}
```

### Using other iterable types

```go
package main

import (
    "fmt"
    it "github.com/elordeiro/go/iters"
)

func main() {
    type IntSlice []int
    slice := it.Iterable(IntSlice{1, 2, 3, 4, 5})
    fmt.Println(slice)
    // Output: ItSlice[int][1 2 3 4 5]

    slice2 := it.Iterable([]string{"a", "b", "c"})
    for i, v := range it.Enumerate(1, slice2) {
        fmt.Printf("%d:%s ", i, v)
    }
    fmt.Println()
    // Output: 1:a 2:b 3:c

    seq1 := it.Iterable([]int{1, 2, 3})
    seq2 := it.Iterable([]string{"a", "b", "c"})
    for v1, v2 := range it.Zip(seq1, seq2) {
        fmt.Printf("%d:%s ", v1, v2)
    }
    fmt.Println()
    // Output: 1:a 2:b 3:c

    slice3 := it.Iterable([]int{0, 1, 2, 3})
    slice4 := it.Iterable([]int{4, 5, 6, 7})
    for v := range it.Chain(slice3, slice4) {
        fmt.Print(v, " ")
    }
    fmt.Println()
    // Output: 0 1 2 3 4 5 6 7

    m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
    isEven := func(k string, v int) bool { return v%2 == 0 }
    for k, v := range it.Filter2(m, isEven) {
        fmt.Printf("%s:%d ", k, v)
    }
    fmt.Println()
    // Output: b:2
}

```

## Notes

-   The `iters` package is still under development and may change in the future.
-   The package is designed to be used with the new iterator support introduced in Go 1.23. It will not work with older versions of Go.
-   The package 'surfaces' the standard library iterators to be used within the package. If you need interact with the standard library, you can use the `Sink()` methods. If you have a standard library iterator and want to use it within the package, you can use the `Surf()` and `Surf2()` functions.

## Contributing

Contributions to the `iters` package are welcome. If you have an idea for a new iterator or an improvement, please open an issue or submit a pull request.

## License

This project is licensed under the BSD-3 License. See the [LICENSE](../LICENSE) file for details.

---
