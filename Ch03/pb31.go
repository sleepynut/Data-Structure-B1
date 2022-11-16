package Ch03

func merge2SortedLists(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}

	if h2 == nil {
		return h1
	}

	var head *ListNode
	if h1.Val < h2.Val {
		head = h1
		h1.Next = merge2SortedLists(h1.Next, h2)
	} else {
		head = h2
		h2.Next = merge2SortedLists(h1, h2.Next)
	}

	return head
}
