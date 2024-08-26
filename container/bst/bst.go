package bst

import (
	"fmt"
	"iter"

	dq "github.com/elordeiro/go/container/deque"
	"golang.org/x/exp/constraints"
)

type Ord constraints.Ordered

// Tree is an AVL tree implementation
type Tree[V Ord] struct {
	Val    V
	Left   *Tree[V]
	Right  *Tree[V]
	Height int
}

// NewBst creates a new AVL tree. If vals are provided, they are added to the tree,
// and the tree is initialized.
func NewBst[V Ord](vals ...V) *Tree[V] {
	var root *Tree[V]
	for _, val := range vals {
		root = root.Insert(val)
	}
	return root
}

// Insert inserts a value into the AVL tree and returns the new root
func (root *Tree[V]) Insert(val V) *Tree[V] {
	if root == nil {
		return &Tree[V]{Val: val, Height: 1}
	}

	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right = root.Right.Insert(val)
	}

	root.Height = 1 + max(height(root.Left), height(root.Right))
	balance := getBalance(root)

	if balance > 1 && val < root.Left.Val {
		return rightRotate(root)
	}

	if balance < -1 && val > root.Right.Val {
		return leftRotate(root)
	}

	if balance > 1 && val > root.Left.Val {
		root.Left = leftRotate(root.Left)
		return rightRotate(root)
	}

	if balance < -1 && val < root.Right.Val {
		root.Right = rightRotate(root.Right)
		return leftRotate(root)
	}

	return root
}

// Delete deletes a value from the AVL tree and returns the new root
func (root *Tree[V]) Delete(val V) *Tree[V] {
	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = root.Left.Delete(val)
	} else if val > root.Val {
		root.Right = root.Right.Delete(val)
	} else {
		if root.Left == nil || root.Right == nil {
			var temp *Tree[V]
			if temp = root.Left; temp == nil {
				temp = root.Right
			}

			if temp == nil {
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := root.Right.min()
			root.Val = temp.Val
			root.Right = root.Right.Delete(temp.Val)
		}
	}

	if root == nil {
		return nil
	}

	root.Height = 1 + max(height(root.Left), height(root.Right))
	balance := getBalance(root)

	if balance > 1 && getBalance(root.Left) >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && getBalance(root.Left) < 0 {
		root.Left = leftRotate(root.Left)
		return rightRotate(root)
	}

	if balance < -1 && getBalance(root.Right) <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && getBalance(root.Right) > 0 {
		root.Right = rightRotate(root.Right)
		return leftRotate(root)
	}

	return root
}

// min returns the minimum value in the tree
func (root *Tree[V]) min() *Tree[V] {
	if root.Left == nil {
		return root
	}
	return root.Left.min()
}

// height returns the height of the tree node
func height[V Ord](node *Tree[V]) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// getBalance returns the balance factor of the tree node
func getBalance[V Ord](root *Tree[V]) int {
	if root == nil {
		return 0
	}

	return height(root.Left) - height(root.Right)
}

// rightRotate performs a right rotation on the tree node
func rightRotate[V Ord](root *Tree[V]) *Tree[V] {
	newRoot := root.Left
	root.Left = newRoot.Right
	newRoot.Right = root

	root.Height = 1 + max(height(root.Left), height(root.Right))
	newRoot.Height = 1 + max(height(newRoot.Left), height(newRoot.Right))

	return newRoot
}

// leftRotate performs a left rotation on the tree node
func leftRotate[V Ord](root *Tree[V]) *Tree[V] {
	newRoot := root.Right
	root.Right = newRoot.Left
	newRoot.Left = root

	root.Height = 1 + max(height(root.Left), height(root.Right))
	newRoot.Height = 1 + max(height(newRoot.Left), height(newRoot.Right))

	return newRoot
}

// ----------------------------------------------------------------------------
// Utils / Traversals
// ----------------------------------------------------------------------------

// Preorder returns an iter.Seq[V] that traverses the tree in preorder
func (root *Tree[V]) Preorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if root == nil {
			return
		}
		yield(root.Val)
		root.Left.Preorder()(yield)
		root.Right.Preorder()(yield)
	}
}

// Inorder returns an iter.Seq[V] that traverses the tree in inorder
func (root *Tree[V]) Inorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if root == nil {
			return
		}
		root.Left.Inorder()(yield)
		yield(root.Val)
		root.Right.Inorder()(yield)
	}
}

// Postorder returns an iter.Seq[V] that traverses the tree in postorder
func (root *Tree[V]) Postorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if root == nil {
			return
		}
		root.Left.Postorder()(yield)
		root.Right.Postorder()(yield)
		yield(root.Val)
	}
}

// Levelorder returns an iter.Seq[V] that traverses the tree in levelorder
func (root *Tree[V]) Levelorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if root == nil {
			return
		}
		q := dq.NewDeque(root)
		level := 0
		for !q.IsEmpty() {
			len := q.Len()
			level++
			for range len {
				node := q.PopFront()
				yield(node.Val)
				if node.Left != nil {
					q.PushBack(node.Left)
				}
				if node.Right != nil {
					q.PushBack(node.Right)
				}
			}
		}
	}
}

// Enumerate returns an iter.Seq2[int, V] that traverses the tree in the given order
func (root *Tree[V]) Enumerate(start int, seq iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		seq(func(val V) bool {
			if !yield(i, val) {
				return false
			}
			i++
			return true
		})
	}
}

// StringOrder returns a string representation of the tree in the specified order
// (preorder, inorder, postorder, levelorder)
func (root *Tree[V]) StringSeq(seq iter.Seq[V]) string {
	s := "/\\["
	seq(func(val V) bool {
		s += fmt.Sprintf("%v ", val)
		return true
	})
	if len(s) > 3 {
		s = s[:len(s)-1]
	}
	return s + "]"
}

// String returns a string representation of the tree in inorder
func (root *Tree[V]) String() string {
	s := "/\\["
	root.Inorder()(func(val V) bool {
		s += fmt.Sprintf("%v ", val)
		return true
	})
	if len(s) > 3 {
		s = s[:len(s)-1]
	}
	return s + "]"
}
