package link

// 编写一个程序，找到两个单链表相交的起始节点。

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	lista, listb := headA, headB
	for lista != listb {
		if lista != nil {
			lista = lista.Next
		}else {
			lista = headB
		}
		if listb != nil {
			listb = listb.Next
		}else {
			listb = headA
		}
	}
	return lista
}
