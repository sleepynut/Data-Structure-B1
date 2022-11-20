package Ch03

func findModularNode(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}

	var matched *ListNode
	idx := 0

	for head != nil {
		idx++
		if idx == k {
			idx = 0
			matched = head
		}
		head = head.Next
	}

	return matched
}
