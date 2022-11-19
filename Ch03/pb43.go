package Ch03

func rmvNodeFromList(head *ListNode, del *ListNode) *ListNode {
	// moving data from next to deleted node wont work
	// if del is the tail itself !!
	var prev *ListNode
	ori := head
	for head != nil {

		if head == del {
			if prev != nil {
				prev.Next = head.Next
				head = ori
			} else {
				head = head.Next
			}
			break
		}

		prev = head
		head = head.Next
	}

	return head
}
