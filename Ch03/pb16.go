package Ch03

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head.Next
	head.Next = nil
	revHead := reverseList(prev)
	prev.Next = head

	return revHead
}
