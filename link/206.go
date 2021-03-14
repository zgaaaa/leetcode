package link

/*
反转一个单链表。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

// type ListNode struct {
//    Val int
//    Next *ListNode
// }

func ReverseList(head *ListNode) *ListNode {
	// 后指针
	var rear *ListNode
	// 错误写法
	// rear := new(ListNode)
	// 因为new初始化会默认为零值,所以反转后的链表末尾会多出一个零
	// 前指针
	front := head
	for front != nil {
		// 中间变量
		next := front.Next
		// 反转,前指针指向后指针
		front.Next = rear
		// 后指针前移
		rear= front
		// 前指针前移
		front = next
		// 以上代码go语言支持更简便写法
		// front, rear, front.Next = front.Next, front, rear 
	}
	return rear
}