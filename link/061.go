package link

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

示例 2:
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/

// type ListNode struct {
//    Val int
//    Next *ListNode
// }

func RotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	// 统计长度
	n := 1
	// 尾指针
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	// 形成闭环
	tail.Next = head
	// 计算

	// k有可能大于n,所以对取余数
	k %= n
	// head和tail向前移动n-k
	for i := 1; i <= n-k; i++ {
		tail = head
		head = head.Next
	}
	// tail指向nil
	tail.Next = nil
	return head
}
