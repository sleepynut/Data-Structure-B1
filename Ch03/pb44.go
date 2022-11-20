package Ch03

func splitEvenOddList(head *ListNode) *ListNode {
	nodes := []*ListNode{}

	// push node to nodes
	h := head
	for h != nil {
		nodes = append(nodes, h)
		h = h.Next
	}

	// move even numbers to the beginning of list
	snowball := 0
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Val%2 != 0 {
			snowball++
			continue
		}
		nodes[i].Val, nodes[i-snowball].Val = nodes[i-snowball].Val, nodes[i].Val
	}

	return head
}
