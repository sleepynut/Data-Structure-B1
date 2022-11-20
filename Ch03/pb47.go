package Ch03

func findModularNodeFromLast(head *ListNode, k int) *ListNode {
	// return the first from last node where nth % k == 0
	_, matched := findFromLastHelper(head, k)
	return matched
}

func findFromLastHelper(head *ListNode, k int) (int, *ListNode) {
	if head == nil || k <= 0 {
		return 0, nil
	}

	pos, matched := findFromLastHelper(head.Next, k)
	pos++
	if pos == k {
		return pos, head
	}

	return pos, matched
}
