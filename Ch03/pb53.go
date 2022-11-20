package Ch03

func addDigitFromTwoLists(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil && h2 == nil {
		return nil
	}

	s1, s2 := twoListsToTwoSlices(h1, h2)

	i1, i2, carry, long := len(s1)-1, len(s2)-1, 0, len(s1)
	if i2 > i1 {
		long = len(s2)
	}

	s3 := make([]*ListNode, long)
	sum := 0
	for i := long - 1; i >= 0; i-- {
		// fmt.Println(i)
		if i1 >= 0 && i2 >= 0 {
			sum = s1[i1].Val + s2[i2].Val + carry
			i1--
			i2--

		} else if i1 < 0 {
			sum = s2[i2].Val + carry
			i2--

		} else {
			sum = s1[i1].Val + carry
			i1--
		}

		s3[i] = &ListNode{
			Val: sum % 10,
		}

		carry = sum / 10

		if i == long-1 {
			s3[i].Next = nil
		} else {
			s3[i].Next = s3[i+1]
		}
	}

	head := s3[0]

	if carry > 0 {
		tmp := &ListNode{
			Val:  carry,
			Next: head,
		}

		head = tmp
	}

	return head
}

func twoListsToTwoSlices(h1 *ListNode, h2 *ListNode) ([]*ListNode, []*ListNode) {
	l1, l2 := []*ListNode{}, []*ListNode{}

	for h1 != nil || h2 != nil {
		if h1 != nil {
			l1 = append(l1, h1)
			h1 = h1.Next
		}

		if h2 != nil {
			l2 = append(l2, h2)
			h2 = h2.Next
		}

	}

	return l1, l2
}
