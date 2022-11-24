package Ch03

func rearrangeFromX(h *ListNode, x int) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}

	dh := &ListNode{Val: -1}
	dt := &ListNode{Val: -2}

	// find tail
	tail := GetTail(h)

	// bind dummy head and tail
	dh.Next = h
	tail.Next = dt

	// moving dummy tail
	dtm := dt

	prev := dh
	for h != dt {

		if h.Val >= x {
			tmp := h.Next
			h.Next = nil

			prev.Next = tmp
			dtm.Next = h

			dtm = dtm.Next
			h = tmp
			continue
		}

		prev = h
		h = h.Next
	}

	if dh.Next == dt {
		return dh.Next.Next
	}

	// free dummy head
	tmp := dh.Next
	dh.Next = nil

	// free dummy tail
	prev.Next = dt.Next
	dt.Next = nil

	dh = tmp
	return dh
}
