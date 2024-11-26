// Package avl provides an AVL tree implementation
package avl

import (
	"fmt"
	"iter"
	"strings"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/containers/deque"
)

// Tree represents a node in the AVL tree.
type Tree[V constraints.Ordered] struct {
	val    V
	left   *Tree[V]
	right  *Tree[V]
	height int
}

// New creates a new AVL tree. If values are provided, they are added to the tree,
// and the tree is initialized.
func New[V constraints.Ordered](vals ...V) *Tree[V] {
	var t *Tree[V]
	for _, val := range vals {
		t = t.Insert(val)
	}
	return t
}

// Value returns the value of the tree node.
func (t *Tree[V]) Value() V {
	if t == nil {
		panic("attempt to dereference nil pointer.\n\tfunc: avl.Value()")
	}
	return t.val
}

// Left returns the left child of the tree node.
func (t *Tree[V]) Left() *Tree[V] {
	if t == nil {
		return nil
	}
	return t.left
}

// Right returns the right child of the tree node.
func (t *Tree[V]) Right() *Tree[V] {
	if t == nil {
		return nil
	}
	return t.right
}

// Height returns the height of the tree node.
func (t *Tree[V]) Height() int {
	if t == nil {
		return 0
	}
	return t.height
}

// Insert inserts a value into the AVL tree and returns the new tree.
func (t *Tree[V]) Insert(val V) *Tree[V] {
	if t == nil {
		return &Tree[V]{val: val, height: 1}
	}

	if val < t.val {
		t.left = t.left.Insert(val)
	} else if val > t.val {
		t.right = t.right.Insert(val)
	} else {
		return t
	}

	t.height = 1 + max(t.left.Height(), t.right.Height())
	balance := t.balance()

	if balance > 1 && val < t.left.val {
		return t.rightRotate()
	}

	if balance < -1 && val > t.right.val {
		return t.leftRotate()
	}

	if balance > 1 && val > t.left.val {
		t.left = t.left.leftRotate()
		return t.rightRotate()
	}

	if balance < -1 && val < t.right.val {
		t.right = t.right.rightRotate()
		return t.leftRotate()
	}

	return t
}

// Delete deletes a value from the AVL tree and returns the new tree.
func (t *Tree[V]) Delete(val V) *Tree[V] {
	if t == nil {
		return nil
	}

	if val < t.val {
		t.left = t.left.Delete(val)
	} else if val > t.val {
		t.right = t.right.Delete(val)
	} else {
		if t.left == nil || t.right == nil {
			var temp *Tree[V]
			if temp = t.left; temp == nil {
				temp = t.right
			}

			if temp == nil {
				t = nil
			} else {
				*t = *temp
			}
		} else {
			temp := t.right.Min()
			t.val = temp.val
			t.right = t.right.Delete(temp.val)
		}
	}

	if t == nil {
		return nil
	}

	t.height = 1 + max(t.left.Height(), t.right.Height())
	balance := t.balance()

	if balance > 1 && t.left.balance() >= 0 {
		return t.rightRotate()
	}

	if balance > 1 && t.left.balance() < 0 {
		t.left = t.left.leftRotate()
		return t.rightRotate()
	}

	if balance < -1 && t.right.balance() <= 0 {
		return t.leftRotate()
	}

	if balance < -1 && t.right.balance() > 0 {
		t.right = t.right.rightRotate()
		return t.leftRotate()
	}

	return t
}

// Search searches for a value in the AVL tree and returns the node that
// contains it if it is found.
func (t *Tree[V]) Search(val V) *Tree[V] {
	if t == nil {
		return nil
	}

	if val < t.val {
		return t.left.Search(val)
	}

	if val > t.val {
		return t.right.Search(val)
	}

	return t
}

// Min returns the node with the minimun value in the tree.
func (t *Tree[V]) Min() *Tree[V] {
	if t == nil {
		return nil
	}
	if t.left == nil {
		return t
	}
	return t.left.Min()
}

// Max returns the node with the  maximum value in the tree.
func (t *Tree[V]) Max() *Tree[V] {
	if t == nil {
		return nil
	}
	if t.right == nil {
		return t
	}
	return t.right.Max()
}

// getBalance returns the balance factor of the tree node.
func (t *Tree[V]) balance() int {
	if t == nil {
		return 0
	}
	return t.left.Height() - t.right.Height()
}

// rightRotate performs a right rotation on the tree node.
func (t *Tree[V]) rightRotate() *Tree[V] {
	newRoot := t.left
	t.left = newRoot.right
	newRoot.right = t

	t.height = 1 + max(t.left.Height(), t.right.Height())
	newRoot.height = 1 + max(newRoot.left.Height(), newRoot.right.Height())

	return newRoot
}

// leftRotate performs a left rotation on the tree node.
func (t *Tree[V]) leftRotate() *Tree[V] {
	newRoot := t.right
	t.right = newRoot.left
	newRoot.left = t

	t.height = 1 + max(t.left.Height(), t.right.Height())
	newRoot.height = 1 + max(newRoot.left.Height(), newRoot.right.Height())

	return newRoot
}

// Preorder returns an iter.Seq[V] that traverses the tree in preorder.
func (t *Tree[V]) Preorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if t == nil {
			return
		}
		yield(t.val)
		t.left.Preorder()(yield)
		t.right.Preorder()(yield)
	}
}

// Inorder returns an iter.Seq[V] that traverses the tree in inorder.
func (t *Tree[V]) Inorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if t == nil {
			return
		}
		t.left.Inorder()(yield)
		yield(t.val)
		t.right.Inorder()(yield)
	}
}

// Postorder returns an iter.Seq[V] that traverses the tree in postorder.
func (t *Tree[V]) Postorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if t == nil {
			return
		}
		t.left.Postorder()(yield)
		t.right.Postorder()(yield)
		yield(t.val)
	}
}

// Levelorder returns an iter.Seq[V] that traverses the tree in levelorder.
func (t *Tree[V]) Levelorder() iter.Seq[V] {
	return func(yield func(V) bool) {
		if t == nil {
			return
		}
		q := deque.New(t)
		for !q.IsEmpty() {
			len := q.Len()
			for range len {
				node := q.PopFront()
				yield(node.val)
				if node.left != nil {
					q.PushBack(node.left)
				}
				if node.right != nil {
					q.PushBack(node.right)
				}
			}
		}
	}
}

// StringOrder returns a string representation of the tree in the specified order.
// The order func can be tree.Preorder, tree.Inorder, tree.Postorder, or tree.Levelorder.
func (t Tree[V]) StringOrder(order func() iter.Seq[V]) string {
	var sb strings.Builder
	sb.WriteString("^[")
	first := true
	for v := range order() {
		if first {
			first = false
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	return sb.String()
}

// String returns a string representation of the tree in inorder.
func (t Tree[V]) String() string {
	var sb strings.Builder
	sb.WriteString("^[")
	first := true
	for v := range t.Inorder() {
		if first {
			first = false
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	return sb.String()
}
