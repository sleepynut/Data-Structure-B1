package Ch03

func findSqrtNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var matched *ListNode
	for i, j := 1, 1; head != nil; i++ {

		if i == j*j {
			if matched != nil {
				matched = matched.Next
			} else {
				matched = head
			}
			j++
		}

		head = head.Next
	}

	return matched
}
