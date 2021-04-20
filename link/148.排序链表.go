package link

import "container/heap"

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/

type PriorityQueue []*ListNode

func (pri PriorityQueue) Len() int            { return len(pri) }
func (pri PriorityQueue) Swap(i, j int)       { pri[i], pri[j] = pri[j], pri[i] }
func (pri PriorityQueue) Less(i, j int) bool  { return pri[i].Val < pri[j].Val }
func (pri *PriorityQueue) Push(x interface{}) { *pri = append(*pri, x.(*ListNode)) }
func (pri *PriorityQueue) Pop() interface{} {
	old := *pri
	n := len(old)
	x := old[n-1]
	*pri = old[0 : n-1]
	return x
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 优先队列
	var h PriorityQueue
	heap.Init(&h)
	// 入队
	for head != nil {
		heap.Push(&h, head)
		head = head.Next
	}
	// 最小的出队
	res := &ListNode{}
	temp := res
	for len(h) > 0 {
		min := heap.Pop(&h).(*ListNode)
		temp.Next = min
		temp = temp.Next
	}
	return res.Next
}
