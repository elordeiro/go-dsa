# Iters Package

> `Go Docs`  
> [![Go Reference](https://pkg.go.dev/badge/github.com/elordeiro/go@v0.0.0-20240819050135-7f0a3c34749a/iters.svg)](https://pkg.go.dev/github.com/elordeiro/go@v0.0.0-20240819050135-7f0a3c34749a/iters)

## Overview

The `iters` package provides a collection of iterators designed to enhance the traversal and transformation of data structures in Go. Leveraging the new iterator support introduced in Go 1.23, this package offers a range of utilities to simplify working with collections.

## Features

-   **Custom Iterators:** Create and use custom iterators for specific needs.
-   **Chainable Operations:** Support for chaining multiple iterator operations like map, filter, and reduce.
-   **Lazy Evaluation:** Efficient processing of elements with lazy evaluation. Process elements only when needed.
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

## Usage

-   Most of the functions in this package return iterators of type `iter.Seq[E]`.
-   If you need that iterator to be of type `iter.Seq2[int, V]`, so that it yields index-value pairs, you can use the `Seq2()` converter to convert it.
-   If you need to start at a different index, you can use the `Enumerate()` function.
-   The `Seq2()` converter is similar to the `Enumerate()` function and defaults to starting at index 0.

### Basic Iterator

Here's an example of using a basic iterator:

```go
package main

import (
    "fmt"
    itr "github.com/elordeiro/go/iters"
)

func main() {
    for v := itr.Range(1, 6){
        fmt.Println(v) // Output: 1 2 3 4 5
    }
}
```

### Chaining Operations

You can chain multiple operations using the iterators:

```go
package main

import (
    "fmt"
    itr "github.com/elordeiro/go/iters"
)

func main() {
    for v := itr.Map(itr.Range(1, 6), func(x int) int { return x * 2 }) {
        fmt.Println(v) // Output: 2 4 6 8 10
    }
}
```

## Iterator Types

-   **Builtin iterators:** The slices / maps package has added support for easily iterating over slices and maps. They both have support for iterating over just values or index-value / key-value pairs.
-   **Modifying iterators:** Once you have an iterator, you can modify it using one of the functions provided in this package.
-   **Range iterators:** Iterate over a range of values, useful for numerical sequences or when you don't need or want a slice just to iterate over it. These include `Range`, `Count`, and `Repeat`.

## Contributing

Contributions to the `iters` package are welcome. If you have an idea for a new iterator or an improvement, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
