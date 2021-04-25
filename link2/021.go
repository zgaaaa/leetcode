package link2

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	result := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Next = l1
			res = res.Next
			l1 = l1.Next
		} else {
			res.Next = l2
			res = res.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		res.Next = l1
	} else {
		res.Next = l2
	}
	return result.Next
}
