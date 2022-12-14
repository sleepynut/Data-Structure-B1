package Ch03

func longerListToSlice(h1 *ListNode, h2 *ListNode) []*ListNode {
	if h1 == nil && h2 == nil {
		return nil
	}
	var slice []*ListNode

	if h1 == nil {
		slice = append(longerListToSlice(nil, h2.Next), h2)
	} else if h2 == nil {
		slice = append(longerListToSlice(h1.Next, nil), h1)
	} else {
		slice = longerListToSlice(h1.Next, h2.Next)
		last := len(slice)

		if slice != nil && h1.Next == slice[last-1] {
			slice = append(slice, h1)
		} else {
			slice = append(slice, h2)
		}
	}
	return slice
}

func findLongerList(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil && h2 == nil {
		return nil
	}

	if h1 == nil {
		return h2
	} else if h2 == nil {
		return h1
	} else {

		prev := findLongerList(h1.Next, h2.Next)
		if h1.Next == prev {
			return h1
		}

		return h2
	}
}

func lengthOfTwoList(h1 *ListNode, h2 *ListNode) (int, int) {
	l1, l2 := 0, 0

	for h1 != nil || h2 != nil {
		if h1 != nil {
			l1++
			h1 = h1.Next
		}

		if h2 != nil {
			l2++
			h2 = h2.Next
		}
	}

	return l1, l2
}

func addDigits(h1 *ListNode, h2 *ListNode, c int) (int, *ListNode) {
	if h1 == nil && h2 == nil {
		return 0, nil
	}

	sum := 0
	res := &ListNode{}

	if h1 == nil {
		sum = (c + h2.Val)
	} else if h2 == nil {
		sum = (c + h1.Val)
	} else {
		sum = (c + h1.Val + h2.Val)
	}

	carry := sum / 10
	res.Val = sum % 10

	return carry, res
}

func addDigitsHelper(l, s *ListNode, skip int) (int, *ListNode) {
	// l: longer ListNode
	// s: shorter ListNode

	// return carry after adding 2 digits and new ListNode
	// New ListNode will contain the added digit

	if l == nil && s == nil {
		return 0, nil
	}

	var next, curr *ListNode
	carry := 0

	if skip > 0 {
		carry, next = addDigitsHelper(l.Next, s, skip-1)
		carry, curr = addDigits(l, nil, carry)
	} else {
		carry, next = addDigitsHelper(l.Next, s.Next, skip)
		carry, curr = addDigits(l, s, carry)
	}

	curr.Next = next
	return carry, curr
}

func addDigitFromTwoLists(h1 *ListNode, h2 *ListNode) *ListNode {
	l1, l2 := lengthOfTwoList(h1, h2)

	long := h1
	short := h2
	skip := l1 - l2

	if l2 > l1 {
		long = h2
		short = h1
		skip = l2 - l1
	}

	carry, head := addDigitsHelper(long, short, skip)

	if carry > 0 {
		tmp := &ListNode{Val: carry, Next: head}
		head = tmp
	}

	return head
}
