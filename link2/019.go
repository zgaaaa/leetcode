package link2

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Next: head}
	l, r := res, res
	pre := res
	i := 1
	for r != nil {
		if i > n {
			pre = l
			l = l.Next
		}
		r = r.Next
		i++
	}
	pre.Next = l.Next
	return res.Next
}
