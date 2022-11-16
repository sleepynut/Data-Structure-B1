package Ch03

func swapPairList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new := head.Next
	moving := new.Next

	head.Next = swapPairList(moving)
	new.Next = head
	head = new

	return head
}
