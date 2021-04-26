package link2

import "container/heap"

type PriorityQueue []*ListNode

func (p PriorityQueue) Len() int            { return len(p) }
func (p PriorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p PriorityQueue) Less(i, j int) bool  { return p[i].Val < p[j].Val }
func (p *PriorityQueue) Push(x interface{}) { *p = append(*p, x.(*ListNode)) }
func (p *PriorityQueue) Pop() interface{} {
	n := *p
	length := len(n)
	x := n[length-1]
	*p = n[:length-1]
	return x
}

func MergeKLists(lists []*ListNode) *ListNode {
	var p PriorityQueue
	for _, v := range lists {
		if v != nil {
			heap.Push(&p, v)
		}
	}
	res := &ListNode{}
	temp := res
	for len(p) > 0 {
		min := heap.Pop(&p).(*ListNode)
		temp.Next = min
		temp = temp.Next
		if min.Next != nil {
			heap.Push(&p, min.Next)
		}
	}
	return res.Next
}
