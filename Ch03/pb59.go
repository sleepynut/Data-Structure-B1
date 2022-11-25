package Ch03

func removeDupFromSortedList(h *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dh := dummy

	var tmp *ListNode
	for h != nil {
		tmp = h.Next

		if dh == dummy || h.Val != dh.Val {
			// detach node
			h.Next = nil

			// bind node to new list
			dh.Next = h
			dh = dh.Next
		}

		h = tmp
	}

	return dummy.Next
}
