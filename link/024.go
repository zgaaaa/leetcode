package link

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

题解:https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/shou-hua-tu-jie-24-liang-liang-jiao-huan-lian-biao/
*/

// type ListNode struct {
//    Val int
//    Next *ListNode
// }

func SwapPairs(head *ListNode) *ListNode {
	// 头指针
	prev := new(ListNode)
	prev.Next = head
	// 用于返回
	result := prev
	// 边界条件不仅是head不能为空,由于需要用到head.Next,所以head.Next也不能为空
	for head != nil && head.Next != nil {
		// 交换
		next := head.Next
		prev.Next, next.Next, head.Next = next, head, next.Next
		// 指针向前移动
		prev, head = head, head.Next
	}
	return result.Next
}
