package link

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{Next: head}
	pre, tail := res, res
	for i := 0; i < right; i++ {
		if i < left-1 {
			pre = pre.Next
		}
		tail = tail.Next
	}
	head = pre.Next
	next := tail.Next
	head, tail = myreverse(head, tail)
	pre.Next = head
	tail.Next = next
	return res.Next
}

func myreverse(head, tail *ListNode) (*ListNode, *ListNode) {
	p := &ListNode{Next: head}
	t := head
	for p != tail {
		p, t, t.Next = t, t.Next, p
	}
	return tail, head
}
