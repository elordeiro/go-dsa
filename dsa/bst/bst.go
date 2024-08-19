package bst

import (
	"math"
	"strconv"

	deque "github.com/elordeiro/go/dsa/deque"
)

var null int = math.MaxInt

// Definition for a binary tree treeNode.
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func NewBinaryTree(nums []int) *treeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	maxIdx := len(nums)
	head := &treeNode{Val: nums[idx]}
	idx++
	queue := deque.NewDeque()
	queue.PushRight(head)

	for !queue.IsEmpty() {
		curr := queue.PopLeft()
		var left int
		var right int
		if idx < maxIdx {
			left = nums[idx]
			idx++
		} else {
			left = null
		}
		if idx < maxIdx {
			right = nums[idx]
			idx++
		} else {
			right = null
		}
		if left != null {
			curr.(*treeNode).Left = &treeNode{Val: left}
			queue.PushRight(curr.(*treeNode).Left)
		}
		if right != null {
			curr.(*treeNode).Right = &treeNode{Val: right}
			queue.PushRight(curr.(*treeNode).Right)
		}
	}
	return head
}

func (root *treeNode) String() string {
	if root == nil {
		return "[]"
	}

	res := "["
	toPrint := deque.NewDeque()
	queue := deque.NewDeque()
	queue.PushRight(root)

	for !queue.IsEmpty() {
		curr := queue.PopLeft()
		if curr.(*treeNode) == nil {
			toPrint.PushRight(null)
			continue
		}
		toPrint.PushRight(curr.(*treeNode).Val)
		queue.PushRight(curr.(*treeNode).Left)
		queue.PushRight(curr.(*treeNode).Right)
	}

	for !toPrint.IsEmpty() && toPrint.PeekRight() == null {
		toPrint.PopRight()
	}

	for !toPrint.IsEmpty() {
		curr := toPrint.PopLeft()
		if curr == null {
			res += "null, "
			continue
		}
		res += strconv.Itoa(curr.(int)) + ", "
	}

	res = res[:len(res)-2]
	res += "]"
	return res
}

func CompareTrees(tree1 *treeNode, tree2 *treeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	if CompareTrees(tree1.Left, tree2.Left) {
		return CompareTrees(tree1.Right, tree2.Right)
	}
	return false
}

func IsSameTree(p *treeNode, q *treeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}
	return false
}

func IsSubtree(root *treeNode, subRoot *treeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if IsSameTree(root, subRoot) {
		return true
	}
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
