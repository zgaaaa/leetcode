package link2

func SwapPairs(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	prev := res
	for head != nil && head.Next != nil {
		next := head.Next
		prev.Next, next.Next, head.Next = next, head, next.Next
		prev = head
		head = prev.Next
	}
	return res.Next
}
