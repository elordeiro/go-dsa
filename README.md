# Go Module

[![Go Reference](https://pkg.go.dev/badge/github.com/elordeiro/go.svg)](https://pkg.go.dev/github.com/elordeiro/go)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](../LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/elordeiro/go)](https://goreportcard.com/report/github.com/elordeiro/go)
[![Release](https://img.shields.io/github/v/release/elordeiro/go)]()

## Overview

This repository contains a collection of go packages that implement some common data structures and algorithms. The library is designed to be modular and easy to use, with each package focusing on a specific data structure, algorithm or some other language feature. The repository also includes iterators that leverage the new Go 1.23 features, along with comprehensive tests for each package.

## Installation

To install this module, run:

```bash
go get github.com/elordeiro/go/
```

You can import specific packages as needed:

```go
import "github.com/elordeiro/go/container/iters"
```

or

```go
import it "github.com/elordeiro/go/container/iters"
```

## Packages

### Container

The [`container`](./container/) collection contains various data structures, each implemented as its own package.

-   [BST](https://pkg.go.dev/github.com/elordeiro/go/container/bst)
-   [Deque](https://pkg.go.dev/github.com/elordeiro/go/container/deque)
-   [List](https://pkg.go.dev/github.com/elordeiro/go/container/list)
-   [PQ (Priority Queue)](https://pkg.go.dev/github.com/elordeiro/go/container/pq)
-   [Set](https://pkg.go.dev/github.com/elordeiro/go/container/set)
-   [Stack](https://pkg.go.dev/github.com/elordeiro/go/container/stack)

### Iters

The [iters](./iters/) package provides helpful iterators that leverage the new Go 1.23 language features. These iterators can be used to traverse or transform collections in a concise and idiomatic manner.

-   **Features:**
    -   Range-based iteration
    -   Filtering and mapping operations
    -   Custom iterators for specific data structures
    -   Lazy evaluation for efficient processing
    -   Much more!

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the BSD-3 License. See the [LICENSE](./LICENSE) file for details.

---
