package Ch03

func findDuplicate(h1 *ListNode, h2 *ListNode) *ListNode {
	// prereqisite: h1 & h2 must be sorted lists

	if h1 == nil || h2 == nil {
		return nil
	}

	var dummy, head *ListNode

	for h1 != nil && h2 != nil {

		if h1.Val < h2.Val {
			h1 = h1.Next
		} else if h2.Val < h1.Val {
			h2 = h2.Next
		} else {

			if head != nil && h1.Val != head.Val {
				head.Next = &ListNode{Val: h1.Val}
				head = head.Next

			} else if head == nil {
				head = &ListNode{Val: h1.Val}
				dummy = head

			}

			h1 = h1.Next
			h2 = h2.Next
		}
	}
	return dummy
}
