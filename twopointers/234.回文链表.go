package twopointers

/*
请判断一个链表是否为回文链表。

示例 1:
输入: 1->2
输出: false

示例 2:
输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一:
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var pre *ListNode
	// 找到mid节点
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 断开
	// 反转
	var head2 *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}
	// 对比
	for head2 != nil && head != nil {
		if head2.Val != head.Val {
			return false
		}
		head2 = head2.Next
		head = head.Next
	}
	return true
}

// 方法二:
func IsPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针遍历链表
	// 同时反转前半链表
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 快指针
		temp := slow.Next // 反转链表
		slow.Next = pre
		pre = slow
		slow = temp
	}
	// 反转完成后,前半段链表的头指针为pre,后半段链表的头指针为slow
	if fast != nil {
		slow = slow.Next // 如果链表为奇数,slow往后挪一个
	}
	// 比较
	for slow != nil && pre != nil {
		if slow.Val != pre.Val {
			return false
		}
		slow = slow.Next
		pre = pre.Next
	}
	return true
}
