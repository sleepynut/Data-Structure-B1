package Ch03

func josephousCircle(n, k int) int {

	if n <= 1 || k <= 0 {
		return -1
	}

	h, prev := GenCircularList(n)

	idx := 1
	for h != h.Next {
		if idx == k {

			// remove kth element
			prev.Next = h.Next

			// reset counting index
			idx = 1

			// move head position
			h = h.Next
		} else {
			idx++
			prev = h
			h = h.Next
		}
	}

	return h.Val
}
