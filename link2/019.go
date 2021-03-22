package link2

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Next: head}
	pre := result
	rear := head
	temp := 0
	for head != nil {
		if temp >= n {
			pre = rear
			rear = rear.Next
		}
		head = head.Next
		temp++
	}
	pre.Next = pre.Next.Next
	return result.Next
}