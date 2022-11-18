package Ch03

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

type RandListNode struct {
	Val  int
	Next *RandListNode
	Rand *RandListNode
}

func GenLinkList(n int) (*ListNode, *ListNode) {
	if n <= 0 {
		return nil, nil
	}

	temp := make([]*ListNode, n)
	for i := 0; i < n; i++ {
		temp[i] = &ListNode{Val: i}
	}
	for i := 0; i < n-1; i++ {
		temp[i].Next = temp[i+1]
	}

	return temp[0], temp[n-1]
}

func GenRandLinkList(n int, rm map[int]int) (*ListNode, *ListNode) {
	if n <= 0 {
		return nil, nil
	}

	temp := make([]*ListNode, n)
	for i := 0; i < n; i++ {
		temp[i] = &ListNode{Val: i}
	}
	for i := 0; i < n-1; i++ {
		temp[i].Next = temp[i+1]
	}

	for k, v := range rm {
		if k >= v || k < 0 || v < 0 || k >= n || v >= n {
			return nil, nil
		}

		temp[k].Next = temp[v]
	}

	return temp[0], temp[n-1]
}

// Provide a circle list of 1,2,3,....,n,1
func GenCircularList(n int) (*ListNode, *ListNode) {
	if n <= 0 {
		return nil, nil
	}

	head, tail := GenLinkList(n + 1)
	tail.Next = head.Next

	// return head, tail of circular list
	return head.Next, tail
}

func GenLinkLisFromValues(nums []int) *ListNode {
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

func PrintList(head *ListNode) string {
	out := ""
	for head != nil {
		out += strconv.Itoa(head.Val) + " "
		head = head.Next
	}
	return out
}

func PrintRandList(head *RandListNode) string {
	if head == nil {
		return ""
	}

	s := strconv.Itoa(head.Val) + " "
	s += PrintRandList(head.Rand)
	s += PrintRandList(head.Next)
	return s
}
