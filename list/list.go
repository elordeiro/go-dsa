package list

import "strconv"

type listNode struct {
	Val  int
	Next *listNode
}

func NewListNode(val int) *listNode {
	node := new(listNode)
	node.Val = val
	return node
}

func CreateList(nums []int) *listNode {
	head := NewListNode(-1)
	curr := head
	for _, num := range nums {
		curr.Next = NewListNode((num))
		curr = curr.Next
	}
	return head.Next
}

func CompareLists(list1 *listNode, list2 *listNode) bool {
	for list1 != nil && list2 != nil && list1.Val == list2.Val {
		list1 = list1.Next
		list2 = list2.Next
	}
	if list1 != nil || list2 != nil {
		return false
	}
	return true
}

func (head *listNode) ToString() string {
	if head == nil {
		return "[]"
	}
	res := "["
	for head != nil {
		res += strconv.Itoa(head.Val) + ", "
		head = head.Next
	}
	res = res[:len(res)-2]
	res += "]"
	return res
}
