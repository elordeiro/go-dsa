[![Go Reference](https://pkg.go.dev/badge/github.com/elordeiro/go/dsa.svg)](https://pkg.go.dev/github.com/elordeiro/go/dsa)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](../LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/elordeiro/go)](https://goreportcard.com/report/github.com/elordeiro/go)
[![Release](https://img.shields.io/github/v/release/elordeiro/go)]()

---

# Container Collection

A collection of generic data structures implemented in Go. This library provides several commonly used data structures such as linked lists, AVL trees, stacks, deques, sets, and priority queues, all implemented with efficiency and ease of use in mind. Each data structure is implemented as a separate package, allowing you to use only the ones you need in your project. All data structures are generic and can store any type of data where possible. Some data structures require the elements to be comparable. All packages contain multiple iterator functions to traverse the data structures efficiently (Go 1.23 or higher).

## Packages

### 1. **list**

The [list](https://pkg.go.dev/github.com/elordeiro/go/container/list) package implements a singly linked list. It allows efficient insertions and deletions from both ends of the list.

**Features:**

-   Add elements at the front or back.
-   Remove single elements or all elements.
-   Retrieve the size of the list.
-   Supports iteration over elements.

### 2. **bst**

The [bst](https://pkg.go.dev/github.com/elordeiro/go/container/bst) package implements a balanced AVL tree. AVL trees maintain their balance through rotations, ensuring operations like insertion, deletion, and search are logarithmic in time complexity.

**Features:**

-   Insert elements while maintaining balance.
-   Delete elements while maintaining balance.
-   Search for elements efficiently.
-   In-order traversal to retrieve sorted elements.
-   Pre-order, post-order and level-order traversal for other use cases.

### 3. **stack**

The [stack](https://pkg.go.dev/github.com/elordeiro/go/container/stack) package implements a simple stack (LIFO) data structure.

**Features:**

-   Push elements onto the stack.
-   Pop elements from the stack.
-   Peek at the top element without removing it.
-   Check if the stack is empty.

### 4. **deque**

The [deque](https://pkg.go.dev/github.com/elordeiro/go/container/deque) package implements a doubly linked list, which allows for efficient insertion and deletion from both ends.

**Features:**

-   Add elements at both the front and the back.
-   Remove elements from both the front and the back.
-   Access elements at either end in constant time.

### 5. **set**

The [set](https://pkg.go.dev/github.com/elordeiro/go/container/set) package implements a set data structure. Sets store unique elements and provide efficient membership checks.

**Features:**

-   Add elements to the set.
-   Remove elements from the set.
-   Check if an element exists in the set.
-   Union, intersection, and difference operations.

### 6. **priority queue**

The [priority queue](https://pkg.go.dev/github.com/elordeiro/go/container/pq) package implements a priority queue data structure, using container/heap for efficient element retrieval based on priority.

**Features:**

-   Builtin Max and Min priority queues.
-   Priority queue with custom priority function.
-   Insert elements with a priority.
-   Extract the element with the highest or lowest priority.

## Getting Started

### Installation

To use the container collection in your Go project, you can get it using `go get`:

```bash
go get github.com/elordeiro/go/container/<package>
```

> For examples, please visit each package's documentation.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue on GitHub.

## License

This project is licensed under the BSD-3 License - see the [LICENSE](../LICENSE) file for details.

---
