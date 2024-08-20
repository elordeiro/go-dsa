# Go Module

> `Go Docs`  
> [![Go Reference](https://pkg.go.dev/badge/github.com/elordeiro/go@v0.0.0-20240819050135-7f0a3c34749a.svg)](https://pkg.go.dev/github.com/elordeiro/go@v0.0.0-20240819050135-7f0a3c34749a)

## Overview

This repository contains a collection of go packages that implement some common data structures and algorithms. The library is designed to be modular and easy to use, with each package focusing on a specific data structure, algorithm or some other language feature. The repository also includes iterators that leverage the new Go 1.23 features, along with comprehensive tests for each package.

## Table of Contents

-   [Installation](#installation)
-   [Packages](#packages)
    -   [DSA](#dsa)
        -   [BST](#bst)
        -   [Deque](#deque)
        -   [List](#list)
        -   [PQ (Priority Queue)](#pq-priority-queue)
        -   [Set](#set)
        -   [Stack](#stack)
    -   [Iters](#iters)
    -   [Tests](#tests)
-   [Usage](#usage)
-   [Contributing](#contributing)
-   [License](#license)

## Installation

To install this module, run:

```bash
go get github.com/elordeiro/go/
```

You can import specific packages as needed:

```go
import "github.com/elordeiro/go/dsa/bst"
```

## Packages

### DSA

The `dsa` collection contains various data structures, each implemented as its own package.

#### BST

The `bst` package provides an implementation of a Binary Search Tree. This data structure supports efficient insertion, deletion, and lookup operations.

-   **Features:**
    -   Insertion of key-value pairs
    -   Deletion of nodes
    -   Search for specific keys
    -   Traversal methods (in-order, pre-order, post-order)

#### Deque

The `deque` package implements a double-ended queue. This data structure allows elements to be added or removed from both ends efficiently.

-   **Features:**
    -   Push and pop operations for both front and back
    -   Peek methods for accessing elements without removal

#### List

The `list` package provides an implementation of a doubly linked list.

-   **Features:**
    -   Insertion and deletion at any position
    -   Iteration over elements
    -   Efficient access to head and tail

#### PQ (Priority Queue)

The `pq` package implements a priority queue using a heap. This data structure is useful for scenarios where you need to process elements based on priority.

-   **Features:**
    -   Enqueue and dequeue based on priority
    -   Peek at the highest priority element

#### Set

The `set` package provides a collection of unique elements with operations such as union, intersection, and difference.

-   **Features:**
    -   Add, remove, and check for elements
    -   Set operations (union, intersection, difference)

#### Stack

The `stack` package implements a LIFO (Last In, First Out) stack.

-   **Features:**
    -   Push and pop operations
    -   Peek at the top element without removal

### Iters

The `iters` package provides helpful iterators that leverage the new Go 1.23 language features. These iterators can be used to traverse or transform collections in a concise and idiomatic manner.

-   **Features:**
    -   Range-based iteration
    -   Filtering and mapping operations
    -   Custom iterators for specific data structures
    -   Lazy evaluation for efficient processing

### Tests

The `tests` package contains comprehensive tests for all the other packages. These tests ensure that the data structures and algorithms behave as expected.

-   **Features:**
    -   Unit tests for each package
    -   Integration tests to verify inter-package compatibility
    -   Test coverage reports

## Usage

Here's an example of how you can use the packages in this module:

```go
package main

import (
    "fmt"
    "github.com/yourusername/dsalib/dsa/bst"
)

func main() {
    tree := bst.New[int, string]()
    tree.Insert(1, "one")
    tree.Insert(2, "two")

    value, found := tree.Search(1)
    if found {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Not found")
    }
}
```

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
