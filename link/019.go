package link

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
进阶：你能尝试使用一趟扫描实现吗？

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

题解：https://www.geekxh.com/1.1.%E9%93%BE%E8%A1%A8%E7%B3%BB%E5%88%97/101.html#_04%E3%80%81%E9%A2%98%E7%9B%AE%E8%A7%A3%E7%AD%94
*/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 头节点
	result := new(ListNode)
	result.Next = head
	// 后指针
	rear := result
	// 指向后指针,方便删除
	pre := new(ListNode)
	// 计算前指针所在的位置
	i := 1
	// 当前指针为空时,表名遍历完毕
	for head != nil {
		// 当前指针的位置大于等于n的时候后指针开始遍历
		if i >= n {
			pre = rear
			rear = rear.Next
		}
		//前指针向后遍历,位置加一
		head = head.Next
		i++
	}
	// 删除节点
	pre.Next = pre.Next.Next
	return result.Next
}
