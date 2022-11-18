package Ch03

func ReverseKthFromList(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}
	idx := 0
	dummy := &ListNode{Val: 1234, Next: head}
	tail := dummy
	head = dummy

	for head != nil {
		idx++
		head = head.Next
		if idx%k == 0 {
			tmp := head.Next
			h, t := reverseHelper(tail.Next, 1, k)

			//bind reversed list
			tail.Next = h
			t.Next = tmp

			// move tail to tail of reversed list
			tail = t

			// restore head <- tmp
			head = tmp

			// since we already move head to head.Next
			// through tmp
			idx++
		}

	}
	return dummy.Next
}

func reverseHelper(head *ListNode, count, end int) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil || count == end {
		return head, head
	}

	newHead, _ := reverseHelper(head.Next, count+1, end)
	head.Next.Next = head
	head.Next = nil

	// return head & tail of the reversed list
	return newHead, head
}
