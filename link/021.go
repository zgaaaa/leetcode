package link

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
*/

// type ListNode struct {
//    Val int
//    Next *ListNode
// }

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 头节点
	head := new(ListNode)
	result := head
	// 循环条件是l1和l2都不能为空,因为为空后无法比较大小
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		}else {
			head.Next = l2
			l2 = l2.Next
		}
		// 头节点向后移动
		head = head.Next
	}
	// 如果l1和l2不等长,会有剩下的节点
	if l1 == nil {
		head.Next = l2
	}
	if l2 == nil {
		head.Next = l1
	}
	return result.Next
}