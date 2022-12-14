package Ch03

func SplitListToTwo(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	var prev *ListNode
	slow := head
	fast := head

	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	// delist
	prev.Next = nil

	return head, slow
}
