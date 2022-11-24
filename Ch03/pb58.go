package Ch03

func rearrangeFromX(h *ListNode, x int) *ListNode {
	dlt, dgt := &ListNode{Val: -1}, &ListNode{Val: -2}
	dmlt, dmgt := dlt, dgt

	var tmp *ListNode
	for h != nil {
		tmp = h.Next

		if h.Val < x {
			dmlt.Next = h
			dmlt = dmlt.Next
		} else {
			dmgt.Next = h
			dmgt = dmgt.Next
		}
		h.Next = nil
		h = tmp
	}

	dmlt.Next = dgt.Next
	return dlt.Next
}
