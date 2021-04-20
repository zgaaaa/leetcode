package link

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，
只保留原始链表中 没有重复出现 的数字。返回同样按升序排列的结果链表。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]
*/

func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{Next: head}
	pre := res
	for head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Val == head.Next.Val {
				head.Next = head.Next.Next
			}
			pre.Next = head.Next
		} else {
			pre = pre.Next
			head = pre.Next
		}
	}

	return res
}
