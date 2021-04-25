package link2

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	hand := res
	temp := 0
	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		hand.Next = &ListNode{Val: temp % 10}
		temp /= 10
		hand = hand.Next
	}
	return res.Next
}
