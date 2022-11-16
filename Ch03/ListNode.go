package Ch03

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func genLinkList(n int) *ListNode {
	if n <= 0 {
		return nil
	}

	temp := make([]*ListNode, n)
	for i := 0; i < n; i++ {
		temp[i] = &ListNode{Val: i}
	}
	for i := 0; i < n-1; i++ {
		temp[i].Next = temp[i+1]
	}

	return temp[0]
}

func genLinkLisFromValues(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	// create default ListNode with nil Next ptr
	temp := make([]*ListNode, n)
	for i := 0; i < n; i++ {
		temp[i] = &ListNode{Val: nums[i]}
	}

	// binding list
	for i := 0; i < n-1; i++ {
		temp[i].Next = temp[i+1]
	}

	return temp[0]
}

func printList(head *ListNode) string {
	out := ""
	for head != nil {
		out += strconv.Itoa(head.Val) + " "
		head = head.Next
	}
	return out
}
