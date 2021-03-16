package link

/*
给定一个链表，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
如果链表中存在环，则返回 true 。否则，返回 false 。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
*/

// type ListNode struct {
// 	Val int
//     Next *ListNode
// }

func HasCycle(head *ListNode) bool {
	// 如果链表没有或只有一个元素,不可能有环
	if head == nil || head.Next == nil {
		return false
	}
    tial := head.Next
	// 如果tial与head相等表名有环,则退出循环返回true
	for tial != head {
		if tial == nil || tial.Next == nil {
			return false
		}
		// 快慢指针遍历
		tial = tial.Next.Next
		head = head.Next
	}
	return true
}