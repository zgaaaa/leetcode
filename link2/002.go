package link2

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	head := result
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		if l1 !=nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{Val: sum%10,Next: nil}
		sum /= 10
		head = head.Next
	}
	return result.Next
}