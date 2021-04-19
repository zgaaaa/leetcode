package link

import "container/heap"

/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]
*/

type PriorityHeap []*ListNode

func (pri PriorityHeap) Len() int            { return len(pri) }
func (pri PriorityHeap) Swap(i, j int)       { pri[i], pri[j] = pri[j], pri[i] }
func (pri PriorityHeap) Less(i, j int) bool  { return pri[i].Val < pri[j].Val }
func (pri *PriorityHeap) Push(x interface{}) { *pri = append(*pri, x.(*ListNode)) }
func (pri *PriorityHeap) Pop() interface{} {
	old := *pri
	n := len(old)
	x := old[n-1]
	*pri = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	var h PriorityHeap
	heap.Init(&h)
	for _, v := range lists {
		if v != nil {
			heap.Push(&h, v)
		}
	}
	res := &ListNode{}
	temp := res
	for len(h) > 0 {
		min := heap.Pop(&h).(*ListNode)
		temp.Next = min
		temp = temp.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return res.Next
}
