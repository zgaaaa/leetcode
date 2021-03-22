package link2

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{Next: head}
	pre := result
	for head != nil {
		// 分组
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return result
			}
		}
		// 反转
		nex := tail.Next
		head, tail = reserve(head,tail)
		// 连接
		pre.Next = head
		tail.Next = nex
		// 指针移动
		pre = tail
		head = nex
	}
	return result.Next
}

func reserve(head, tail *ListNode) (*ListNode, *ListNode) {
	h, p := head, &ListNode{Next: head}
	for p != tail {
		h.Next, h, p = p, h.Next, h
	}
	return tail, head
}
