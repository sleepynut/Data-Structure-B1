package Ch03

import (
	"strconv"
	"testing"
)

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

func walkList(head *ListNode) string {
	out := ""
	for head != nil {
		out += strconv.Itoa(head.Val)
		head = head.Next
	}
	return out
}

func TestReverseList(t *testing.T) {
	size := 5
	head := genLinkList(size)

	rev := reverseList(head)

	expect := "43210"
	get := walkList(rev)

	if get != expect {
		t.Errorf("Expect: %s BUT get: %s\n", expect, get)
	}
}
