package Ch03

func rearrangeNodes(h *ListNode) *ListNode {
	l := GetLength(h)
	mid := l / 2

	if l < 3 {
		return h
	}

	ptr1, _ := rearrangeHelper(h, mid, l)
	return ptr1
}

func rearrangeHelper(h *ListNode, stop int, l int) (*ListNode, *ListNode) {
	// return ptr1 & ptr2 to represent moving PrintList
	// ptr1: 1,2,3,4,...
	// ptr2: n,n-1,n-2,...

	if stop == 0 {
		if l%2 == 0 {
			return nil, h
		} else {
			tmp := h.Next

			// detach mid element
			h.Next = nil

			return h, tmp
		}
	}

	ptr1, ptr2 := rearrangeHelper(h.Next, stop-1, l)

	h.Next = ptr2

	tmp := ptr2.Next
	ptr2.Next = ptr1

	// move ptr2 to match returned stack of h
	ptr2 = tmp

	return h, ptr2
}
